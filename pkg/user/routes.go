package user

import (
    "github.com/go-chi/chi/v5"
    "instagram-bis/config"
)

func Routes(cfg *config.Config) chi.Router {
    r := chi.NewRouter()

    r.Post("/users/register", RegisterUser(cfg))
    r.Post("/users/login", LoginUser(cfg))
    r.Get("/users/{id}", GetUserProfile(cfg))
    r.Put("/users/{id}", UpdateUserProfile(cfg))
    r.Post("/users/{id}/follow", FollowUser(cfg))
    r.Get("/users/{id}/followers", GetFollowers(cfg))
    r.Get("/users/{id}/following", GetFollowing(cfg))

    return r
}