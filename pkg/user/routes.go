package user

import (
	"instagram-bis/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()

	r.Post("/register", RegisterUser(cfg))
	r.Post("/login", LoginUser(cfg))
	r.Get("/{id}", GetUserProfile(cfg))
	r.Put("/{id}", UpdateUserProfile(cfg))
	r.Post("/{id}/follow", FollowUser(cfg))
	r.Get("/{id}/followers", GetFollowers(cfg))
	r.Get("/{id}/following", GetFollowing(cfg))

	return r
}
