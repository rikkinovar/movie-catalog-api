package handlers

import (
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse :nodoc
type ErrorResponse struct {
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

// Render :nodoc
func (response ErrorResponse) Render(writer http.ResponseWriter, request *http.Request) error {
	render.Status(request, response.StatusCode)
	return nil
}

var errBadRequest = ErrorResponse{
	StatusCode: 400,
	Message:    "Bad request",
}

var errUnauthorized = ErrorResponse{
	StatusCode: 401,
	Message:    "Unauthorized",
}

var errUnprocessableEntity = ErrorResponse{
	StatusCode: 422,
	Message:    "Unprocessable entity",
}

var errForbidden = ErrorResponse{
	StatusCode: 403,
	Message:    "Forbidden",
}

var errNotFound = ErrorResponse{
	StatusCode: 404,
	Message:    "Resource not found",
}

var errInternalServerError = ErrorResponse{
	StatusCode: 500,
	Message:    "Internal server error",
}
