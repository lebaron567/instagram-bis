package authentication

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/login", LoginHandler)

	return r
}
