package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No environment sevrer file found")
	}
}

func main() {
	connString := os.Getenv("DB_CONN_STRING")

	db := sqlx.MustConnect("postgres", connString)
	var version string
	err := db.QueryRow("SELECT version()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(`Database version: `, version)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true, // Enable debugging for testing, consider disabling in production
	})

	const user = "chris"

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested the %s's user data", user)
	})

	mux.HandleFunc("/db-version", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The database  version is: %s\n", version)
	})

	// Wrap the handler with the CORS middleware
	handler := corsMiddleware.Handler(mux)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", handler)
}
