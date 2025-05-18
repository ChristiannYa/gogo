package routes

import (
	"fmt"
	"net/http"
)

// UserHandler handles user-related requests
func UserHandler(w http.ResponseWriter, r *http.Request) {
	const user = "chris"
	fmt.Fprintf(w, "Hello, you've requested the %s's user data", user)
}
