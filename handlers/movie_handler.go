package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/rikkinovar/movie-catalog-api/handlers/requests"
	"github.com/rikkinovar/movie-catalog-api/handlers/responses"
	"github.com/rikkinovar/movie-catalog-api/models"
	"github.com/rikkinovar/movie-catalog-api/services"
)

type MovieHandler struct {
	movieService services.MovieServiceContract
}

//GetRoutes return routes
func (handler *MovieHandler) GetRoutes() chi.Router {
	router := chi.NewRouter()

	router.Get("/", handler.Index)
	router.Get("/{id}", handler.Find)
	router.Post("/", handler.Create)
	router.Put("/{id}", handler.Update)
	router.Delete("/{id}", handler.Delete)

	return router
}

// Index :nodoc
// @Summary Return all movies
// @Description Return all movies
// @Tags movies
// @Produce json
// @Accept json
// @Success 200 {object} responses.MovieListResponse
// @Failure 500 {object} ErrorResponse
// @Router /movies [get]
func (handler *MovieHandler) Index(writer http.ResponseWriter, request *http.Request) {
	movieList, err := handler.movieService.Get()
	if err != nil {
		render.Render(writer, request, errInternalServerError)
		return
	}

	response := responses.CreateMovieListResponse(movieList)
	err = render.Render(writer, request, response)
	if err != nil {
		render.Render(writer, request, errInternalServerError)
	}
}

// Index :nodoc
// @Summary Return movie by ID
// @Description Return moview by ID
// @Tags movies
// @Produce json
// @Accept json
// @Success 200 {object} responses.MovieResponse
// @Failure 500 {object} ErrorResponse
// @Router /movies [get]
func (handler *MovieHandler) Find(writer http.ResponseWriter, request *http.Request) {
	paramID := chi.URLParam(request, "id")
	id, _ := strconv.Atoi(paramID)
	movie, err := handler.movieService.Find(id)
	if err != nil {
		render.Render(writer, request, errNotFound)
	}

	response := responses.CreateMovieResponse(movie)
	err = render.Render(writer, request, response)
	if err != nil {
		render.Render(writer, request, errInternalServerError)
	}
}

// Index :nodoc
// @Summary Insert Movie
// @Description Insert Movie
// @Tags movies
// @Produce json
// @Accept json
// @Success 200 {object} responses.MovieResponse
// @Failure 500 {object} ErrorResponse
// @Router /movies [post]
func (handler *MovieHandler) Create(writer http.ResponseWriter, request *http.Request) {
	requestData := requests.CreateMovieRequests{}
	if err := render.Bind(request, &requestData); err != nil {
		errResponse := errBadRequest
		if err.Error() != "EOF" {
			errResponse.Message = err.Error()
		}

		render.Render(writer, request, errResponse)
		return
	}

	releaseDate, _ := time.Parse("2006-01-02", requestData.ReleaseDate)
	movie, err := handler.movieService.Create(models.Movie{
		Title:       requestData.Title,
		Genre:       requestData.Genre,
		ImdbRating:  requestData.ImdbRating,
		ReleaseDate: releaseDate,
		PosterUrl:   requestData.PosterUrl,
		Synopsis:    requestData.Synopsis,
	})
	if err != nil {
		render.Render(writer, request, errInternalServerError)
		return
	}

	response := responses.CreateMovieResponse(movie)
	err = render.Render(writer, request, response)
	if err != nil {
		render.Render(writer, request, errInternalServerError)
		return
	}
}

// Index :nodoc
// @Summary Update Movie
// @Description Update Movie
// @Tags movies
// @Produce json
// @Accept json
// @Success 200 {object} responses.MovieResponse
// @Failure 500 {object} ErrorResponse
// @Router /movies/{id} [put]
func (handler *MovieHandler) Update(writer http.ResponseWriter, request *http.Request) {
	var requestData requests.CreateMovieRequests

	idParam := chi.URLParam(request, "id")
	id, _ := strconv.Atoi(idParam)

	if err := render.Bind(request, &requestData); err != nil {
		errResponse := errUnprocessableEntity
		if err.Error() != "EOF" {
			errResponse.Message = err.Error()
		}

		render.Render(writer, request, errResponse)
		return
	}

	releaseDate, _ := time.Parse("2006-01-02", requestData.ReleaseDate)
	movie, err := handler.movieService.Update(id, models.Movie{
		Title:       requestData.Title,
		Genre:       requestData.Genre,
		ImdbRating:  requestData.ImdbRating,
		ReleaseDate: releaseDate,
		PosterUrl:   requestData.PosterUrl,
		Synopsis:    requestData.Synopsis,
	})
	if err != nil {
		render.Render(writer, request, errBadRequest)
		return
	}

	response := responses.CreateMovieResponse(movie)
	if err = render.Render(writer, request, response); err != nil {
		render.Render(writer, request, errInternalServerError)
		return
	}
}

// Index :nodoc
// @Summary Delete Movie
// @Description Delete Movie
// @Tags movies
// @Produce json
// @Accept json
// @Success 200 {object} responses.MovieResponse
// @Failure 500 {object} ErrorResponse
// @Router /movies/{id} [delete]
func (handler *MovieHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	idParam := chi.URLParam(request, "id")
	id, _ := strconv.Atoi(idParam)

	movie, err := handler.movieService.Find(id)
	if err != nil {
		render.Render(writer, request, errNotFound)
		return
	}

	success, err := handler.movieService.Delete(id)
	if err != nil || !success {
		render.Render(writer, request, errInternalServerError)
		return
	}

	if err := render.Render(writer, request, responses.CreateMovieResponse(movie)); err != nil {
		render.Render(writer, request, errInternalServerError)
		return
	}
}
