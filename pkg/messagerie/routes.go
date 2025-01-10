package messagerie

import (
	"instagram-bis/config"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes enregistre les routes du message
func RegisterRoutes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Post("/", CreateMessage(cfg))                         // Créer un message
	r.Get("/discussion/{id}", GetMessagesByDiscussion(cfg)) // Récupérer les messages d'une discussion
	r.Put("/{id}", UpdateMessage(cfg))                      // Mettre à jour un message
	r.Delete("/{id}", DeleteMessage(cfg))                   // Supprimer un message
	return r
}
