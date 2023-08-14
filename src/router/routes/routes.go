package routes

import (
	"api/src/middlewares"
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
	routes = append(routes, loginRoute)

	for _, route := range routes {

		if route.RequireAuth {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)

		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
