package handler

import (
	"github.com/labstack/echo/v4"
	"go-echo-movies/internal/service"
	"net/http"
)

// MovieController handles HTTP requests related to movies
type MovieController struct {
	movieService *service.MovieService
}

// NewMovieController creates a new MovieController
func NewMovieController(movieService *service.MovieService) *MovieController {
	return &MovieController{movieService: movieService}
}

// GetMovies is a handler that returns a list of movies
//func (mc *MovieController) GetMovies(c echo.Context) error {
//	movies := mc.movieService.GetMovies()
//	return c.JSON(http.StatusOK, movies)
//}

// GetMovies retrieves all movies
func (h *MovieController) GetMovies(c echo.Context) error {

	movies, err := h.movieService.GetMovies()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to retrieve movies"})
	}

	return c.JSON(http.StatusOK, movies)
}
