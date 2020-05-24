package middleware

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
	"time"
)

type Middleware interface {
	ServeHTTP(http.ResponseWriter, *http.Request, http.HandlerFunc)
}

type log struct {
	logger *logrus.Logger
}

func NewLoggerMiddleware(l *logrus.Logger ) Middleware {
	return &log{logger: l}
}

func (l *log) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	t := time.Now()
	l.logger.WithFields(getRequestFields(r)).Info("Request")
	next(w, r)

	fields := getResponseFields(r, w.(negroni.ResponseWriter))
	fields["Duration"] = time.Since(t)
	l.logger.WithFields(fields).Info("Response")
}

func getRequestFields(r *http.Request) logrus.Fields {
	fields := logrus.Fields{}
	fields["Client"] = r.RemoteAddr
	fields["Method"] = r.Method
	fields["URL"] = r.URL.String()
	fields["Referrer"] = r.Referer()
	fields["User-Agent"] = r.UserAgent()

	return fields
}

func getResponseFields(r *http.Request, w negroni.ResponseWriter) logrus.Fields {
	fields := logrus.Fields{}
	fields["Method"] = r.Method
	fields["URL"] = r.URL.String()
	fields["StatusCode"] = w.Status()
	fields["Size"] = w.Size()

	return fields
}