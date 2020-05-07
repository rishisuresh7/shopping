package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"shopping-service/constants"
)

type Success struct {
	Result interface{}	`json:"result"`
}

type Error struct {
	Error interface{} `json:"error"`
}

func (s Success) Send(w http.ResponseWriter) error {
	data, err := json.Marshal(s)
	if err != nil {
		fmt.Errorf("unable to encode interface")
	}
	setHeaders(w, data, constants.Success)

	return nil
}

func (e Error) ServerError(w http.ResponseWriter) error {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("unable to encode interface")
	}

	setHeaders(w, data, constants.InternalServerError)
	return nil
}

func (e Error) ClientError(w http.ResponseWriter) error {
	data, err := json.Marshal(e)
	if err != nil {
		return fmt.Errorf("unable to encode interface")
	}

	setHeaders(w, data, constants.ClientError)
	return nil
}

func setHeaders(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	_,_ = w.Write(data)
}