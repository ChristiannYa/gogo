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

	"go-intro/server/routes"
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

	mux := routes.SetupRoutes(version)

	// Wrap the handler with the CORS middleware
	handler := corsMiddleware.Handler(mux)

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", handler)
}
