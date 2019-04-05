package main

import (
	"fast-track/handlers"
	"fast-track/models"

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
	e.Logger.Fatal(e.Start(":9000"))
}

func init() {
	models.Init()
}
