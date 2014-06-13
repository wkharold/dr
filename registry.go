package dr

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Context struct {
	Telemetry bool
}

type Registry struct {
	*mux.Router
}

func New(ctx Context) (*Registry, error) {
	r := mux.NewRouter()

	if ctx.Telemetry {
		r.Handle("/debug/vars", http.DefaultServeMux)
	}

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
