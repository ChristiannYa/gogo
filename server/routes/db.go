package routes

import (
	"fmt"
	"net/http"
)

// DBVersionHandler returns the database version
func DBVersionHandler(dbVersion string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The database version is: %s\n", dbVersion)
	}
}
