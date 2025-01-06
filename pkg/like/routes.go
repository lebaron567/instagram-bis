package like

import (
	"instagram-bis/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Post("/{id}/like", LikePost(cfg))
	r.Delete("/{id}/like", UnlikePost(cfg))

	return r
}
