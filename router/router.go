package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"shopping-service/constants"
	"shopping-service/factory"
	"shopping-service/handler"
)

//Router ....router for shopping service
func Router(f factory.Factory, l *logrus.Logger) *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/login", handler.Login(f, l)).Methods("POST")
	router.HandleFunc("/healthcheck", handler.HealthCheck(f, l)).Methods(constants.GET)
	router.HandleFunc("/{type}/getcollection", handler.GetCollection(f, l)).Methods(constants.GET)
	router.HandleFunc("/{type}/getcategories", handler.GetCategories(f, l)).Methods(constants.GET)

	return router
}