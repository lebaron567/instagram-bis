package post

import (
	"instagram-bis/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Post("/", CreatePost(cfg))
	r.Get("/{id}", GetPost(cfg))
	r.Delete("/{id}", DeletePost(cfg))
	r.Get("/user/{id}", GetPostsByUser(cfg))
	r.Get("/feed", GetAllPosts(cfg))

	return r
}
