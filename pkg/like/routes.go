package like

import (
    "github.com/go-chi/chi/v5"
    "instagram-bis/config"
)

func Routes(cfg *config.Config) chi.Router {
    r := chi.NewRouter()

    r.Post("/posts/{id}/like", LikePost(cfg))
    r.Delete("/posts/{id}/like", UnlikePost(cfg))

    return r
}