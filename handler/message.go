package handler

import (
	"net/http"

	"golang-web-api-hands-on/handler/response"
	"golang-web-api-hands-on/usecase"

	"github.com/go-chi/render"
)

type Message interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type messageHandler struct {
	useCase usecase.Message
}

func NewMessage(useCase usecase.Message) Message {
	return &messageHandler{
		useCase: useCase,
	}
}

func (m *messageHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	message := m.useCase.Get(ctx)

	if message == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, response.ErrNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"message": message,
	})
}
