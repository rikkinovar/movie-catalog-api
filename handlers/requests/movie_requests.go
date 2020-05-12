package requests

import (
	"errors"
	"net/http"
	"time"
)

type CreateMovieRequests struct {
	Title       string  `json:"title"`
	Genre       string  `json:"genre"`
	ImdbRating  float64 `json:"imdb_rating"`
	ReleaseDate string  `json:"release_date"`
	PosterUrl   string  `json:"poster_url"`
}

func (req CreateMovieRequests) Bind(request *http.Request) error {
	if req.Title == "" {
		return errors.New("`title` is requred")
	}
	if req.Genre == "" {
		return errors.New("`genre` is requred")
	}
	_, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		return errors.New("invalid format for `releaseDate`")
	}
	if req.ReleaseDate == "" {
		return errors.New("`releaseDate` is requred")
	}
	if req.PosterUrl == "" {
		return errors.New("`posterUrl` is requred")
	}
	return nil
}
