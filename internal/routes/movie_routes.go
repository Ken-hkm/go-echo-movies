package routes

import (
	"github.com/labstack/echo/v4"
	"go-echo-movies/internal/handler"
	"go-echo-movies/internal/service"
	"gorm.io/gorm"
)

// RegisterMovieRoutes sets up all user-related routes
func RegisterMovieRoutes(e *echo.Echo, db *gorm.DB) {
	movieService := service.NewMovieService(db) // Pass the initialized db
	movieController := handler.NewMovieController(movieService)

	// Define user-related routes here
	e.GET("/movie", movieController.GetMovies)
	//e.POST("/user", userController.CreateUser)
}
