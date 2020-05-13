package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/rikkinovar/movie-catalog-api/handlers"
	httpSwagger "github.com/swaggo/http-swagger"
)

var version = "1.0.0"

func createRouter() chi.Router {
	router := chi.NewRouter()
	swaggerURL := fmt.Sprintf("http://localhost:%s/swagger/doc.json", os.Getenv("APP_PORT"))

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(render.SetContentType(render.ContentTypeJSON))
	router.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(swaggerURL)))

	router.Get("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		payload := map[string]interface{}{
			"name":    "Movie catalog API",
			"version": version,
		}
		response, _ := json.Marshal(payload)
		writer.Write(response)
	}))

	router.Mount("/movies", (handlers.NewMovieHandler(movieService)).GetRoutes())

	return router
}

func serveHTTP() {
	// router := createRouter()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}

	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	fmt.Printf("App running on port %s\n", port)
}
