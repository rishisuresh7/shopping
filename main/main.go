package main

import (
	"fmt"
	"shopping-service/config"

	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"

	"shopping-service/factory"
	"shopping-service/router"
)

func main() {

	c := &config.AppConfig{
		MongoDatabase: "shoppingDB",
		MongoPort: 6000,
		MongoHost: "localhost",
		MongoPassword: "admin",
		MongoUsername: "admin",
	}

	logger := logrus.New()
	f := factory.NewFactory(c, logger)
	r := router.Router(f, logger)

	n := negroni.New()
	n.Use(f.NewLoggerMiddleware())
	n.UseHandler(r)
	n.Run(fmt.Sprintf(":%d", 4000))
}