package authentication

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) http.Handler {
	r := chi.NewRouter()

	r.Post("/login", LoginHandler(db))

	return r
}
