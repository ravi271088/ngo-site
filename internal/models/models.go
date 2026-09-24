package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title" validate:"required,min=3,max=100"`
	Description string    `json:"description" validate:"required"`
	EventDate   time.Time `json:"event_date" validate:"required"`
	Location    string    `json:"location" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

type GalleryImage struct {
	ID         uuid.UUID `json:"id"`
	ImageURL   string    `json:"image_url" validate:"required,url"`
	Caption    string    `json:"caption"`
	Category   string    `json:"category" validate:"required"`
	UploadedAt time.Time `json:"uploaded_at"`
}
