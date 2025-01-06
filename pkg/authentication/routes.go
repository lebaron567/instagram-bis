package authentication

import (
	"instagram-bis/config"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) http.Handler {
	r := chi.NewRouter()
	r.Post("/login", LoginHandler(cfg))
	return r
}
