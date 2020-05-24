package handler

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"shopping-service/factory"
	"shopping-service/models"
	"shopping-service/response"
)

func Login(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.UserCred
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			l.Errorf("Login: unable to decode payload: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		if len(user.Username) < 5 || len(user.Password) < 5 {
			l.Errorf("Login: invalid username or password")
			response.Error{Error: "invalid request"}.ClientError(w)
			return
		}

		u := f.NewUser("users")
		token, err := u.Login(user.Username, user.Password)
		if err != nil {
			l.Errorf("Login: unable to login: %s", err)
			response.Error{Error: "unexpected error happened"}.ServerError(w)
			return
		}

		w.Header().Set("Authorization", token)
		response.Success{Result: "logged in successfully"}.Send(w)
	}
}