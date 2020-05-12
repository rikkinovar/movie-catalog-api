package responses

import (
	"net/http"

	"github.com/rikkinovar/movie-catalog-api/models"
)

type MovieResponse struct {
	Data models.Movie `json:"data"`
}

type MovieListResponse struct {
	Data []models.Movie `json:"data"`
}

// Render :nodoc
func (response MovieResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

// Render :nodoc
func (response MovieListResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

// CreateMovieListResponse returns new MovieListResponse.
func CreateMovieListResponse(movieList []models.Movie) MovieListResponse {
	return MovieListResponse{
		Data: movieList,
	}
}

// CreateMovieResponse returns new MovieResponse.
func CreateMovieResponse(movie models.Movie) MovieResponse {
	return MovieResponse{
		Data: movie,
	}
}
