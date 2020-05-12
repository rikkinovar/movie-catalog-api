package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/rikkinovar/movie-catalog-api/repositories"
	"github.com/rikkinovar/movie-catalog-api/services"
)

var dbMovieRepository repositories.MovieRepositoryContract

var movieService services.MovieServiceContract

func initRepositories(dbConn *gorm.DB, httpClient *http.Client) {
	dbMovieRepository = repositories.NewDBMovieRepository(dbConn)
}

func initServices() {
	movieService = services.NewMovieService(dbMovieRepository)
}
