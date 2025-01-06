package main

import (
	
	"log"
	"net/http"
	"os"

	"instagram-bis/config"
	"instagram-bis/pkg/authentication"
	"instagram-bis/pkg/comment"
	"instagram-bis/pkg/conversation"
	"instagram-bis/pkg/like"
	"instagram-bis/pkg/messagerie"
	"instagram-bis/pkg/post"
	"instagram-bis/pkg/user"

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
		r.Mount("/auth", authentication.Routes(cfg))
		r.Mount("/comment", comment.Routes(cfg))
		r.Mount("/users", user.Routes(cfg))
		r.Mount("/like", like.Routes(cfg))
		r.Mount("/discussions", conversation.RegisterRoutes(cfg))
		r.Mount("/messages", messagerie.RegisterRoutes(cfg))
		r.Mount("/posts", post.Routes(cfg))
	})

	

	log.Printf("Serveur démarré sur le port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
