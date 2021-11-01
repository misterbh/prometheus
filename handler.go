package prometheus

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/portainer/portainer/api/http/security"
)

// Handler is the HTTP handler used to handle PROMETHEUS operations.
type Handler struct {
	*mux.Router
}

// NewHandler returns a new Handler
func NewHandler(bouncer *security.RequestBouncer) *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}
	h.Handle("/prometheus",
		bouncer.RestrictedAccess(http.HandlerFunc(h.prometheus))).Methods(http.MethodGet)

	return h
}
