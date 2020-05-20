package driver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"shopping-service/models"
	"time"
)

type MongoConnector interface {
	Reader(bson.D, string) (interface{}, error)
	Writer() error
}

type mongoDriverConnector struct {
	client *mongo.Client
	database string
}

func NewMongoDriver(c *mongo.Client, d string) MongoConnector{
	return &mongoDriverConnector{client: c, database: d}
}

func (m *mongoDriverConnector) Reader(filter bson.D, collection string) (interface{}, error) {
	var result interface{}
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	session,err := m.client.StartSession()
	if err != nil {
		return nil, fmt.Errorf("unable to create session: %s", err)
	}
	defer session.EndSession(ctx)

	coll := m.client.Database(m.database).Collection(collection)
	err = mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		cursor, err := coll.Find(ctx, filter)
		if err != nil {
			return err
		}
		result, err = m.parse(cursor, sc)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("unable to connect to database")
	}

	return result, nil
}

func (m *mongoDriverConnector) Writer() error {
	return nil
}

func (m *mongoDriverConnector) parse(cursor *mongo.Cursor, c mongo.SessionContext) (interface{}, error) {
	var res models.Result
	for cursor.Next(c) {
		err := cursor.Decode(&res)
		if err != nil {
			return nil, fmt.Errorf("%s",err)
		}

	}

	return res, nil
}