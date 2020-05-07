package healthcheck

import (
	"io"
	"net/http"
)
//Application App details
type Application interface {
	HealthCheck(w http.ResponseWriter)
	Version()
}

type applicationDetails struct {

}
//NewApplicationDetails app details object
func NewApplicationDetails() Application {
	return &applicationDetails {
	}
}

func (a *applicationDetails) HealthCheck(w http.ResponseWriter) {
	io.WriteString(w, "Iam Alive")
}

func (a *applicationDetails) Version() {
	
}