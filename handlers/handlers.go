package handlers

import "github.com/rikkinovar/movie-catalog-api/services"

func NewMovieHandler(movieService services.MovieServiceContract) *MovieHandler {
	return &MovieHandler{
		movieService: movieService,
	}
}
