package handler

import (
	"net/http"
	"shopping-service/factory"
	"shopping-service/response"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func GetCollection(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cartType, ok := vars["type"]
		if !ok {
			l.Errorf("GetCollection: could not read 'type' from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		collectionDriver := f.NewCollection(cartType)
		data, err := collectionDriver.GetCollection()

		if err != nil {
			l.Errorf("GetCollection: unable to get collection")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Result: data}.Send(w)
	}
}

func GetCategories(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		cartType, ok := vars["type"]
		if !ok {
			l.Errorf("GetCollection: could not read 'type' from path params")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		collectionDriver := f.NewCollection(cartType)
		data, err := collectionDriver.GetCategories()

		if err != nil {
			l.Errorf("GetCollection: unable to get collection")
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		response.Success{Result: data}.Send(w)
	}
}
