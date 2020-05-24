package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	formatter "github.com/x-cray/logrus-prefixed-formatter"

	"shopping-service/config"
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
		ApiKey: "foobarbaz",
	}

	exclude := []string{"/healthcheck", "/login"}
	logger := logrus.New()
	logger.Formatter = &formatter.TextFormatter{
		ForceColors: false,
		ForceFormatting: true,
		FullTimestamp: true,
		TimestampFormat: "2006 Jan 02 03:04:05 PM",
	}
	f := factory.NewFactory(c, logger, "foobarbaz")
	r := router.Router(f, logger)

	n := negroni.New()
	n.Use(f.NewLoggerMiddleware())
	n.Use(f.NewAuthenticationMiddleware(exclude, c.ApiKey))
	n.UseHandler(r)
	n.Run(fmt.Sprintf(":%d", 4000))
}