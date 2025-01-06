package conversation

import (
	"github.com/go-chi/chi/v5"
)

func Routes() {
	r := chi.NewRouter()

	r.Get("/", GetConversations)
	r.Get("/{id}", GetConversation)
	r.Post("/", CreateConversation)
	r.Post("/{id}/messages", CreateMessage)
	r.Delete("/{id}", DeleteConversation)

	return r
}
