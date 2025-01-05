package comment

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "instagram-bis/config"
)

type Comment struct {
    ID      uint   `gorm:"primaryKey"`
    PostID  uint   `gorm:"index"`
    Content string `gorm:"type:text"`
}

func AddComment(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid post ID", http.StatusBadRequest)
            return
        }

        var comment Comment
        if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        comment.PostID = uint(postID)
        if err := cfg.DB.Create(&comment).Error; err != nil {
            http.Error(w, "Failed to add comment", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(comment)
    }
}

func GetComments(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid post ID", http.StatusBadRequest)
            return
        }

        var comments []Comment
        if err := cfg.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
            http.Error(w, "Failed to get comments", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(comments)
    }
}

func DeleteComment(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid comment ID", http.StatusBadRequest)
            return
        }

        if err := cfg.DB.Delete(&Comment{}, commentID).Error; err != nil {
            http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}