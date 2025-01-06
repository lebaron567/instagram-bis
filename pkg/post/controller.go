package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	"instagram-bis/config"
	"instagram-bis/database/dbmodel"

	"github.com/go-chi/chi/v5"
)


func CreatePost(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var post dbmodel.Post
		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		if _, err := cfg.PostRepository.Create(&post); err != nil {
			http.Error(w, "Failed to create post", http.StatusInternalServerError)
			return
		}

		reponse := dbmodel.Post{
			ID:      post.ID,
			IDUser:  post.IDUser,
			Content: post.Content,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(reponse)
	}
}


func GetPost(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		post, err := cfg.PostRepository.FindByID(postID)
		if err != nil {
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(post)
	}
}


func DeletePost(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		postID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid post ID", http.StatusBadRequest)
			return
		}

		if err := cfg.PostRepository.Delete(postID); err != nil {
			http.Error(w, "Failed to delete post", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}


func GetPostsByUser(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		posts, err := cfg.PostRepository.FindByUserID(userID)
		if err != nil {
			http.Error(w, "Failed to get posts", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
	}
}


func GetAllPosts(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := cfg.PostRepository.FindAll()
		if err != nil {
			http.Error(w, "Failed to get posts", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
	}
}
