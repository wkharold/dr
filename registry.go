package dr

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// Registry is implemented by types in the dr package that respond to Docker Registry
// API requests.
type Registry interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// Context collects up all the context required by a DockerRegistry instance.
type Context struct {
	Telemetry bool
	AccessOut io.Writer
}

// DockerRegistry instances are responsible for responding to Docker Registry API requests.
type DockerRegistry struct {
	*mux.Router
}

// New uses the given Context to instantiate a type that implements Registry and so be used
// to respond to Docker Registry API requests.
func New(ctx Context) (Registry, error) {
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

	if ctx.AccessOut != nil {
		return &AccessLogger{DockerRegistry{r}, ctx.AccessOut}, nil
	}

	return &DockerRegistry{r}, nil
}
