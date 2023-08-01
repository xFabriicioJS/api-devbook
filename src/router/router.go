package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Run is responsible for starting the router
func Run() *mux.Router {
	r := mux.NewRouter()

	routes.SetupRoutes(r)

	return r
}
