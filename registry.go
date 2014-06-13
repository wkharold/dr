package dr

import "github.com/gorilla/mux"

type Registry struct {
	*mux.Router
}

func New() (*Registry, error) {
	r := mux.NewRouter()

	v1 := r.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/_ping", ping)
	v1.HandleFunc("/images/{imageid}/layer", layer)
	v1.HandleFunc("/images/{imageid}/json", json)
	v1.HandleFunc("/images/{imageid}/ancestry", ancestry)
	v1.HandleFunc("/repositories/{namespace}/{repository}/tags", tags)
	v1.HandleFunc("/repositories/{namespace}/{repository}/tags/{tag}", tag)
	v1.HandleFunc("/repositories/{namespace}/{repository}/", repository)

	return &Registry{r}, nil
}
