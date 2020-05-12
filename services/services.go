package services

import "github.com/rikkinovar/movie-catalog-api/repositories"

func NewMovieService(dbMovieRepository repositories.MovieRepositoryContract) MovieServiceContract {
	return &MovieService{
		dbMovieRepository: dbMovieRepository,
	}
}
