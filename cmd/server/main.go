package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"ngo-site/internal/api/handler"
	"ngo-site/internal/api/middleware"
	"ngo-site/internal/repository"

	_ "github.com/lib/pq"
)

// renderTemplate is a helper to render a page using the base layout
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// We parse both the base layout and the specific page template
	files := []string{
		"web/templates/layouts/base.html",
		filepath.Join("web/templates/pages", tmpl+".html"),
	}

	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Template error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Execute the "base.html" template, which will include the "content" block
	err = t.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		log.Printf("Execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

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
		renderTemplate(w, "home", map[string]interface{}{"Title": "Home"})
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "about", map[string]interface{}{"Title": "About Us"})
	})

	mux.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "contact", map[string]interface{}{"Title": "Contact Us"})
	})

	mux.HandleFunc("/donate", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "donate", map[string]interface{}{"Title": "Donate Now"})
	})

	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "admin_dashboard", map[string]interface{}{"Title": "Admin Dashboard"})
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
