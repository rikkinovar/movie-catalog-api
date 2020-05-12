package services

import "github.com/rikkinovar/movie-catalog-api/models"

type MovieServiceContract interface {
	Get() ([]models.Movie, error)
	Find(id int) (models.Movie, error)
	Create(data models.Movie) (models.Movie, error)
	Delete(id int) (bool, error)
	Update(id int, data models.Movie) (models.Movie, error)
}
