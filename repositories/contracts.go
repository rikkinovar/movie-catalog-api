package repositories

import "github.com/rikkinovar/movie-catalog-api/models"

type MovieRepositoryContract interface {
	Get() ([]models.Movie, error)
	Find(id int) (models.Movie, error)
	Create(models.Movie) (models.Movie, error)
	Delete(id int) (bool, error)
	Update(id int, data models.Movie) (models.Movie, error)
}
