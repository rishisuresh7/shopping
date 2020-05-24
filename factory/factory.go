package factory

import (
	"context"
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"shopping-service/authentication"
	"shopping-service/collection"
	"shopping-service/config"
	"shopping-service/driver"
	"shopping-service/healthcheck"
	"shopping-service/middleware"
	"shopping-service/user"
)
//Factory Object for all methods
type Factory interface {
	NewApplicationDetails() healthcheck.Application
	NewExecuter() driver.MongoConnector
	NewCollection(string) collection.Collection
	NewLoggerMiddleware() middleware.Middleware
	NewAuthenticationMiddleware([]string, string) middleware.Middleware
	NewUser(coll string) user.User
}

var mongoConn sync.Once

type factory struct {
	mongoConn *mongo.Client
	config *config.AppConfig
	logger *logrus.Logger
	secret string
}

//NewFactory Object Initialisation
func NewFactory(c *config.AppConfig, l *logrus.Logger, secret string) Factory{
	return &factory {
		config: c,
		logger : l,
		secret: secret,
	}
}

func (f *factory) NewMongoDriver() (*mongo.Client, error) {
	var (
		connErr error
	)
	mongoConn.Do(func () {
		clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d",
			f.config.MongoUsername,
			f.config.MongoPassword,
			f.config.MongoHost,
			f.config.MongoPort))
		c, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			connErr = err
			return
		}

		f.mongoConn = c
		connErr = err
	})

	if connErr != nil {
		return nil, fmt.Errorf("unable to connect to MongoDB")
	}

	return f.mongoConn, nil
}

func (f *factory) NewExecuter() driver.MongoConnector {
	client, err := f.NewMongoDriver()
	if err != nil || client == nil {
		f.logger.WithError(err).Fatalf("Factory: unable to connect to mongoDB")
	}

	return driver.NewMongoDriver(client, f.config.MongoDatabase)
}

func (f *factory) NewCollection(coll string) collection.Collection {
	return collection.NewCollection(f.NewExecuter(), coll)
}

func (f *factory) NewApplicationDetails() healthcheck.Application {
	return healthcheck.NewApplicationDetails()
}

func (f *factory) NewLoggerMiddleware() middleware.Middleware {
	return middleware.NewLoggerMiddleware(f.logger)
}

func (f *factory) NewAuthenticationMiddleware(r []string, token string) middleware.Middleware {
	return authentication.NewAuthenticationMiddleware(f.logger, r, token, f.secret)
}

func (f *factory) NewUser(coll string) user.User {
	return user.NewUser(coll, f.secret, f.NewExecuter())
}