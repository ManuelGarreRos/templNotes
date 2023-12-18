package main

import (
	"github.com/ManuelGarreRos/templNotes/handler"
	"github.com/labstack/echo/v4"
)

type DB struct{}

func main() {
	app := echo.New()

	// Initialize handler
	noteHandler := handler.NoteHandler{}
	// Register routes
	app.GET("/notes", noteHandler.GetNotes)

	err := app.Start(":3000")
	if err != nil {
		return
	}
}
