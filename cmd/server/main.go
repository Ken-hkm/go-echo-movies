package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-echo-movies/internal/db"
	internalMiddleware "go-echo-movies/internal/middleware"
	"go-echo-movies/internal/routes"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	//DB Connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PWD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	// Connect to the database
	dbConn, err := db.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	// Run migrations
	db.MigrateDB(dbConn)

	//init echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(internalMiddleware.AuthMiddleware) // Custom middleware

	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	// Register routes
	routes.RegisterMovieRoutes(e, dbConn)

	e.Logger.Fatal(e.Start(":8080"))
}
