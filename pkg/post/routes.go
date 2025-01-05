package post

import (
    "github.com/go-chi/chi/v5"
    "instagram-bis/config"
)

func Routes(cfg *config.Config) chi.Router {
    r := chi.NewRouter()

    r.Post("/posts", CreatePost(cfg))
    r.Get("/posts/{id}", GetPost(cfg))
    r.Delete("/posts/{id}", DeletePost(cfg))
    r.Get("/posts/user/{id}", GetPostsByUser(cfg))
    r.Get("/posts/feed", GetAllPosts(cfg))

    return r
}