package routes

import (
	"net/http"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(dbVersion string) *http.ServeMux {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/user", UserHandler)
	mux.HandleFunc("/db-version", DBVersionHandler(dbVersion))

	return mux
}
