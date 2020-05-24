package authentication

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"

	"shopping-service/constants"
	"shopping-service/middleware"
	"shopping-service/response"
	"shopping-service/utilities"
)

type authenticator struct {
	routes []string
	logger *logrus.Logger
	apiKey  string
	secret string
}

func NewAuthenticationMiddleware(l *logrus.Logger, r []string, t string, s string) middleware.Middleware {
	return &authenticator{logger: l, routes: r, apiKey: t, secret: s}
}

func (a *authenticator) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	if utilities.CheckList(r.RequestURI, a.routes) {
		next(w, r)
		return
	}

	if r.Header.Get("api-key") != a.apiKey {
		a.logger.Errorf("Authenticator: forbidden")
		response.Error{Error: "forbidden"}.Forbidden(w)
		return
	}

	token := r.Header.Get("Authorization")
	if ok, err := a.verifyToken(token); !ok {
		a.logger.Errorf("Authenticator: %s", err)
		response.Error{Error: "unauthorized"}.UnAuthorized(w)
		return
	}

	next(w, r)
}

func (a *authenticator) verifyToken(t string) (bool, error) {
	tokenString := strings.Split(t, constants.Bearer)
	if len(tokenString) != 2{
		return false, fmt.Errorf("invalid token")
	}

	newClaims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString[1], newClaims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(a.secret), nil
	})

	if  err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}