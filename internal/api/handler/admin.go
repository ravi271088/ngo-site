package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ngo-site/internal/models"
	"ngo-site/internal/repository"
	"time"

	"github.com/google/uuid"
)

type AdminHandler struct {
	repo *repository.PostgresRepository
}

func NewAdminHandler(repo *repository.PostgresRepository) *AdminHandler {
	return &AdminHandler{repo: repo}
}

// Public API
func (h *AdminHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := h.repo.GetEvents()
	if err != nil {
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func (h *AdminHandler) GetGallery(w http.ResponseWriter, r *http.Request) {
	images, err := h.repo.GetGallery()
	if err != nil {
		http.Error(w, "Failed to fetch gallery", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

// Admin API
func (h *AdminHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		EventDate   string `json:"event_date"`
		Location    string `json:"location"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	parsedDate, err := time.Parse("2006-01-02T15:04", input.EventDate)
	if err != nil {
		http.Error(w, "Invalid date format. Use YYYY-MM-DDTHH:MM", http.StatusBadRequest)
		return
	}

	event := models.Event{
		ID:          uuid.New(),
		Title:       input.Title,
		Description: input.Description,
		EventDate:   parsedDate,
		Location:    input.Location,
		CreatedAt:   time.Now(),
	}

	if err := h.repo.CreateEvent(&event); err != nil {
		http.Error(w, fmt.Sprintf("Failed to create event: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}

func (h *AdminHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteEvent(id); err != nil {
		http.Error(w, "Failed to delete event", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *AdminHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var img models.GalleryImage
	if err := json.NewDecoder(r.Body).Decode(&img); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	img.ID = uuid.New()
	if err := h.repo.CreateGalleryImage(&img); err != nil {
		http.Error(w, "Failed to upload image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(img)
}

func (h *AdminHandler) DeleteImage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.repo.DeleteGalleryImage(id); err != nil {
		http.Error(w, "Failed to delete image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
