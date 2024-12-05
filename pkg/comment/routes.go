package comment


import (
    "github.com/go-chi/chi/v5"
    "net/http"
)

func Routes() chi.Router {
    router := chi.NewRouter()
    router.Post("/posts/:id/comments",r)
    router.Get("/posts/:id/comments", )
	router.Delete("/comments/:id,", )
    return router
}

