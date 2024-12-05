package post

import "github.com/go-chi/chi"

func post() chi.Router {
	r := chi.NewRouter()
	r.Post("/posts", createPost)
	r.Get("/posts/{id}", getPostId)
	r.Delete("/posts/{id}", deletePost)
	r.Get("/posts/user/{id}", deletePostUser)
	r.Get("/posts/feed", getAllPosts)
	return r
}
