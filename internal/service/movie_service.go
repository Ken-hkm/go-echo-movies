package service

import (
	"go-echo-movies/internal/model"
	"gorm.io/gorm"
)

type MovieService struct {
	db *gorm.DB
}

// NewMovieService creates a new MovieService instance with a valid *gorm.DB that returns *MovieService
func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{
		db: db,
	}
}

// GetMovies retrieves all movies from the database
func (s *MovieService) GetMovies() ([]model.Movie, error) {
	var movies []model.Movie
	err := s.db.Find(&movies).Error
	return movies, err
}
