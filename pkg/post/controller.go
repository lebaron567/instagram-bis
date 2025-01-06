package post

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "instagram-bis/config"
    "instagram-bis/database/dbmodel"
)

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body dbmodel.Post true "Post"
// @Success 201 {object} dbmodel.Post
// @Failure 400 {string} string "Invalid request"
// @Failure 500 {string} string "Failed to create post"
// @Router /posts [post]
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

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(post)
    }
}

// GetPost godoc
// @Summary Get a post by ID
// @Description Get a post by ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} dbmodel.Post
// @Failure 400 {string} string "Invalid post ID"
// @Failure 404 {string} string "Post not found"
// @Router /posts/{id} [get]
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

// DeletePost godoc
// @Summary Delete a post by ID
// @Description Delete a post by ID
// @Tags posts
// @Param id path int true "Post ID"
// @Success 204
// @Failure 400 {string} string "Invalid post ID"
// @Failure 500 {string} string "Failed to delete post"
// @Router /posts/{id} [delete]
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

// GetPostsByUser godoc
// @Summary Get posts by user ID
// @Description Get posts by user ID
// @Tags posts
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} dbmodel.Post
// @Failure 400 {string} string "Invalid user ID"
// @Failure 500 {string} string "Failed to get posts"
// @Router /posts/user/{id} [get]
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

// GetAllPosts godoc
// @Summary Get all posts
// @Description Get all posts
// @Tags posts
// @Produce json
// @Success 200 {array} dbmodel.Post
// @Failure 500 {string} string "Failed to get posts"
// @Router /posts/feed [get]
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