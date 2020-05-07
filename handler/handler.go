package handler

import (
	"shopping-service/factory"
	"net/http"
	"github.com/sirupsen/logrus"
)

//HealthCheck handler for healthCheck
func HealthCheck(f factory.Factory, l *logrus.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		healthCheck := f.NewApplicationDetails()
		healthCheck.HealthCheck(w)
	}
}