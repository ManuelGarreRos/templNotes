package handler

import (
	"github.com/ManuelGarreRos/templNotes/view/note"
	"github.com/labstack/echo/v4"
)

type NoteHandler struct{}

func (n *NoteHandler) GetNotes(c echo.Context) error {
	return render(c, note.ShowNotes())
}
