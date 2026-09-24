package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"ngo-site/internal/api/handler"
	"ngo-site/internal/api/middleware"
	"ngo-site/internal/repository"

	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Database Connection
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	defer db.Close()

	// Initialize Repository & Schema
	repo := repository.NewPostgresRepository(db)
	if err := repo.InitSchema(); err != nil {
		log.Fatalf("Error initializing schema: %s", err)
	}

	// Handlers
	adminH := handler.NewAdminHandler(repo)

	mux := http.NewServeMux()

	// Public Routes
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Server is healthy")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the NGO Website!")
	})

	// Admin API Routes (Protected)
	mux.HandleFunc("/api/admin/events", middleware.AdminAuth(adminH.CreateEvent))
	mux.HandleFunc("/api/admin/events/delete", middleware.AdminAuth(adminH.DeleteEvent))
	mux.HandleFunc("/api/admin/gallery", middleware.AdminAuth(adminH.UploadImage))

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
