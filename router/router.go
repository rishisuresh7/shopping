package router

import (
	"shopping-service/handler"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"shopping-service/factory"
)

//Router ....router for shopping service
func Router(f factory.Factory, l *logrus.Logger) *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/healthcheck", handler.HealthCheck(f, l))
	router.HandleFunc("/{type}/getcollection", handler.GetCollection(f, l))
	router.HandleFunc("/{type}/getcategories", handler.GetCategories(f, l))

	return router
}