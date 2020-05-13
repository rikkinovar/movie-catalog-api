package models

import "time"

type Movie struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Genre       string     `json:"genre"`
	ImdbRating  float64    `json:"imdb_rating"`
	ReleaseDate time.Time  `json:"release_date"`
	Synopsis    string     `json:"synopsis"`
	PosterUrl   string     `json:"poster_url"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"type:decimal(10)" json:"deleted_at"`
}
