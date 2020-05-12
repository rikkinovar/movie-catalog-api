package services

import (
	"github.com/rikkinovar/movie-catalog-api/models"
	"github.com/rikkinovar/movie-catalog-api/repositories"
)

type MovieService struct {
	dbMovieRepository repositories.MovieRepositoryContract
}

//Get returns all movies
func (service *MovieService) Get() ([]models.Movie, error) {
	return service.dbMovieRepository.Get()
}

// Find return moview by ID
func (service *MovieService) Find(id int) (models.Movie, error) {
	return service.dbMovieRepository.Find(id)
}

// Create save new movie
func (service *MovieService) Create(data models.Movie) (models.Movie, error) {
	return service.dbMovieRepository.Create(data)
}

//Delete soft delete movie by ID
func (service *MovieService) Delete(id int) (bool, error) {
	return service.dbMovieRepository.Delete(id)
}

//Update movie
func (service *MovieService) Update(id int, data models.Movie) (models.Movie, error) {
	return service.dbMovieRepository.Update(id, data)
}
