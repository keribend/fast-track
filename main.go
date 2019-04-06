package main

import (
	"fast-track/handlers"
	"fast-track/models"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Define the HTTP routes
	e.Static("/public", "public")
	e.GET("/questionnaire", handlers.GetQuestionnaire())
	e.POST("/questionnaire", handlers.AnswerQuestionnaire())
	e.File("/", "public/index.html")

	// Start server

	port := os.Getenv("PORT")

	if port == "" {
		e.Logger.Fatal("$PORT must be set")
	}

	e.Logger.Fatal(e.Start(":" + port))
}

func init() {
	models.Init()
}
