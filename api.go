package dr

import "net/http"

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Vary", "Accept")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Docker-Registry-Version", "0.6.0")
	w.WriteHeader(200)
}
