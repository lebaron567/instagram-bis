package main

import (
	"fmt"
	"instagram-bis/config"
	"instagram-bis/pkg/authentication"
	"instagram-bis/pkg/comment"
	"instagram-bis/pkg/conversation"
	"instagram-bis/pkg/like"
	"instagram-bis/pkg/messagerie"
	"instagram-bis/pkg/post"
	"instagram-bis/pkg/user"
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
		r.Mount("/auth", authentication.Routes())
		r.Mount("/", user.Routes(cfg))
		r.Mount("/", like.Routes(cfg))
		r.Mount("/", conversation.RegisterRoutes(cfg))
		r.Mount("/", messagerie.RegisterRoutes(cfg))
		r.Mount("/", post.Routes(cfg))
	})

	r.Group(func(r chi.Router) {
		r.Use(authentication.AuthMiddleware("c8f9d72e3b4a6d9e7f0b1c2a3e4f5g6h7i8j9k0l1m2n3o4p5q6r7s8t9u0v1w2x3"))

		r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
			user := authentication.GetUserFromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("Welcome, %s!", user)))
		})
	})

	log.Printf("Serveur démarré sur le port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Erreur lors du démarrage du serveur : %v", err)
	}
}
