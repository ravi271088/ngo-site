package repository

import (
	"database/sql"
	"ngo-site/internal/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Event methods
func (r *PostgresRepository) CreateEvent(e *models.Event) error {
	query := `INSERT INTO events (id, title, description, event_date, location, created_at)
              VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(query, e.ID, e.Title, e.Description, e.EventDate, e.Location, e.CreatedAt)
	return err
}

func (r *PostgresRepository) GetEvents() ([]models.Event, error) {
	rows, err := r.db.Query(`SELECT id, title, description, event_date, location, created_at FROM events ORDER BY event_date ASC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var e models.Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Description, &e.EventDate, &e.Location, &e.CreatedAt); err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}

func (r *PostgresRepository) DeleteEvent(id uuid.UUID) error {
	_, err := r.db.Exec(`DELETE FROM events WHERE id = $1`, id)
	return err
}

// Gallery methods
func (r *PostgresRepository) CreateGalleryImage(img *models.GalleryImage) error {
	query := `INSERT INTO gallery (id, image_url, caption, category, uploaded_at)
              VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, img.ID, img.ImageURL, img.Caption, img.Category, img.UploadedAt)
	return err
}

func (r *PostgresRepository) GetGallery() ([]models.GalleryImage, error) {
	rows, err := r.db.Query(`SELECT id, image_url, caption, category, uploaded_at FROM gallery ORDER BY uploaded_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []models.GalleryImage
	for rows.Next() {
		var img models.GalleryImage
		if err := rows.Scan(&img.ID, &img.ImageURL, &img.Caption, &img.Category, &img.UploadedAt); err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, nil
}

func (r *PostgresRepository) DeleteGalleryImage(id uuid.UUID) error {
	_, err := r.db.Exec(`DELETE FROM gallery WHERE id = $1`, id)
	return err
}

func (r *PostgresRepository) InitSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS events (
		id UUID PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		description TEXT,
		event_date TIMESTAMP NOT NULL,
		location VARCHAR(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS gallery (
		id UUID PRIMARY KEY,
		image_url VARCHAR(255) NOT NULL,
		caption VARCHAR(255),
		category VARCHAR(50),
		uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := r.db.Exec(schema)
	return err
}
