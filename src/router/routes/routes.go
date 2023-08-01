package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

// SetupRoutes is responsible for setting up the routes in the router, it receives a router and returns a router
func SetupRoutes(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
