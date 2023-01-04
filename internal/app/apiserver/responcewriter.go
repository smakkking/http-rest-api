package apiserver

import "net/http"

type responseWriter struct {
	http.ResponseWriter // anonymous field
	code                int
}

func (w *responseWriter) WriteHeader(statusCode int) {
	w.code = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
