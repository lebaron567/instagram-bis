package comment

import (
    "github.com/go-chi/chi/v5"
    "instagram-bis/config"
)

func Routes(cfg *config.Config) chi.Router  {
    r := chi.NewRouter()

    r.Post("/posts/{id}/comments", AddComment(cfg))
    r.Get("/posts/{id}/comments", GetComments(cfg))
    r.Delete("/{id}", DeleteComment(cfg))
	return r
}