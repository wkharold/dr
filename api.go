package dr

import (
	_ "expvar"
	"net/http"
)

const (
	drv = "0.7.0"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Vary", "Accept")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Docker-Registry-Version", drv)
	w.WriteHeader(200)
}

func layer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func json(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func ancestry(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func tags(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func tag(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func repository(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}
