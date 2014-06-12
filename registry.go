package dr

import "github.com/gorilla/mux"

type Registry struct {
	*mux.Router
}

func New() (*Registry, error) {
	r := mux.NewRouter()
	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/_ping", ping)
	return &Registry{r}, nil
}
