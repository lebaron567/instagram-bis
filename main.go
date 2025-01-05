package main

import (
	"instagram-bis/config"
	"instagram-bis/pkg/comment"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Erreur lors de l'initialisation de la configuration : %v", err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/api/v1", func(r chi.Router) {
		r.Mount("/comment", comment.Routes(cfg))
	})

	log.Printf("Serveur démarré sur le port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
