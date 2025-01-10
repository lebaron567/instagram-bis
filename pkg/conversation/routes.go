package conversation

import (
	"instagram-bis/config"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes enregistre toutes les routes pour la gestion des discussions
func RegisterRoutes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Post("/", CreateDiscussion(cfg))                   // Créer une nouvelle discussion
	r.Get("/user/{userID}", GetDiscussionsByUserID(cfg)) // Récupérer toutes les discussions pour un utilisateur
	r.Get("/{id}", GetDiscussionByID(cfg))               // Récupérer une discussion par ID
	r.Delete("/{id}", DeleteDiscussion(cfg))             // Supprimer une discussion par ID
	return r
}
