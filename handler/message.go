package handler

import (
	"golang-web-api-hands-on/handler/response"
	"golang-web-api-hands-on/usecase"

	"github.com/go-chi/chi/v5"
)

type Message interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type messageHandler struct {
	useCase usecase.Message
}

func NewMessageHandler(useCase usecase.Message) Message {
	return &messageHandler{
		useCase: useCase,
	}
}

func (m *messageHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	messageID := chi.URLParam(r, "messageID")

	if messageID == "" {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, response.ErrNotFound)
	}

	result, err := b.useCase.Get(ctx, messageID)
	if err != nil {
		if err.Error() == "could not find message" {
			render.Status(r, http.StatusNotFound)
			render.JSON(w, r, response.ErrNotFound)
			return
		}
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, response.ErrSystemError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, response.NewMessage(result))
}
