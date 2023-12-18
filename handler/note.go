package handler

import (
	"github.com/ManuelGarreRos/templNotes/model"
	"github.com/ManuelGarreRos/templNotes/view/note"
	"github.com/labstack/echo/v4"
)

type NoteHandler struct{}

func (n *NoteHandler) GetNotes(c echo.Context) error {
	newNote := []model.Note{
		{
			ID:    1,
			Title: "My first note",
			Body:  "This is the body of my first note",
		},
		{
			ID:    2,
			Title: "My second note",
			Body:  "This is the body of my second note",
		},
		{
			ID:    3,
			Title: "My third note",
			Body:  "This is the body of my third note",
		},
	}

	return render(c, note.ShowNotes(newNote))
}
