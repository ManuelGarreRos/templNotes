package handler

import (
	"github.com/ManuelGarreRos/templNotes/model"
	"github.com/ManuelGarreRos/templNotes/repositories"
	"github.com/ManuelGarreRos/templNotes/view/note"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
)

func NewNoteHandler(log *zap.Logger, nr *repositories.NoteRepository) *NoteHandler {
	return &NoteHandler{
		log:            log,
		NoteRepository: nr,
	}
}

type NoteHandler struct {
	log            *zap.Logger
	NoteRepository *repositories.NoteRepository
}

func (n *NoteHandler) GetNotes(c echo.Context) error {
	notes, err := n.NoteRepository.GetNotes()
	if err != nil {
		return err
	}

	return render(c, note.ShowNotes(notes))
}

func (n *NoteHandler) CreateNote(c echo.Context) error {
	newNote := &model.Note{
		ID:    uuid.New().String(),
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
	}

	err := n.NoteRepository.CreateNote(newNote)
	if err != nil {
		return err
	}

	http.Redirect(c.Response(), c.Request(), "/notes", http.StatusSeeOther)
	return nil
}
