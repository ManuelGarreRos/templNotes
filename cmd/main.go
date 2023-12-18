package main

import (
	"github.com/ManuelGarreRos/templNotes/appctr"
	"github.com/ManuelGarreRos/templNotes/handler"
	"github.com/ManuelGarreRos/templNotes/migrations"
	"github.com/ManuelGarreRos/templNotes/repositories"
	"github.com/labstack/echo/v4"
)

func main() {
	appctr.Start()
	app := echo.New()
	db := appctr.DB()
	log := appctr.Log()

	migrations.MigrateSchema(db)
	// Initialize Repositories
	noteRepositorie := repositories.NewNoteRepository(db, log)
	// Initialize handler
	noteHandler := handler.NewNoteHandler(log, noteRepositorie)
	// Register routes
	app.GET("/notes", noteHandler.GetNotes)
	app.POST("/create-note", noteHandler.CreateNote)

	err := app.Start(":3000")
	if err != nil {
		return
	}
}
