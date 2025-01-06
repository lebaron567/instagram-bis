package messagerie

import (
	"instagram-bis/config"

	"github.com/go-chi/chi/v5"
)

// RegisterRoutes enregistre les routes du message
func RegisterRoutes(r chi.Router, cfg *config.Config) {
	r.Post("/messages", CreateMessage(cfg))                          // Créer un message
	r.Get("/messages/discussion/{id}", GetMessagesByDiscussion(cfg)) // Récupérer les messages d'une discussion
	r.Put("/messages/{id}", UpdateMessage(cfg))                      // Mettre à jour un message
	r.Delete("/messages/{id}", DeleteMessage(cfg))                   // Supprimer un message
}
