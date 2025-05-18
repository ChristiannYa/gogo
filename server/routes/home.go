package routes

import (
	"fmt"
	"net/http"
)

// HomeHandler handles the root path
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}
