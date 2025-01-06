package comment

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "instagram-bis/config"
    "instagram-bis/database/dbmodel"
    model "instagram-bis/pkg/models"
)


func AddComment(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid post ID", http.StatusBadRequest)
            return
        }

        var comment dbmodel.Comment
        if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        comment.IDPost = uint(postID)
        if _, err := cfg.CommentRepository.Create(&comment); err != nil {
            http.Error(w, "Failed to add comment", http.StatusInternalServerError)
            return
        }

        response := model.Comment{
            ID:      comment.ID,
            IDUser:  comment.IDUser,
            IDPost:  comment.IDPost,
            Content: comment.Content,
        }
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(response)
    }
}


func GetComments(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid post ID", http.StatusBadRequest)
            return
        }

        comments, err := cfg.CommentRepository.FindByPostID(postID)
        if err != nil {
            http.Error(w, "Failed to get comments", http.StatusInternalServerError)
            return
        }

        var response []model.Comment
        for _, comment := range comments {
            response = append(response, model.Comment{
                ID:      comment.ID,
                IDUser:  comment.IDUser,
                IDPost:  comment.IDPost,
                Content: comment.Content,
            })
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(response)
    }
}


func DeleteComment(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        commentID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid comment ID", http.StatusBadRequest)
            return
        }

        if err := cfg.CommentRepository.Delete(commentID); err != nil {
            http.Error(w, "Failed to delete comment", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}