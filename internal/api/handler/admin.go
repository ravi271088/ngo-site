package handler

import (
	"encoding/json"
	"net/http"
	"ngo-site/internal/models"
	"ngo-site/internal/repository"

	"github.com/google/uuid"
)

type AdminHandler struct {
	repo *repository.PostgresRepository
}

func NewAdminHandler(repo *repository.PostgresRepository) *AdminHandler {
	return &AdminHandler{repo: repo}
}

func (h *AdminHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	event.ID = uuid.New()
	if err := h.repo.CreateEvent(&event); err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
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
