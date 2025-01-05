package like

import (
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "instagram-bis/config"
    "instagram-bis/database/dbmodel"
)

func LikePost(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid post ID", http.StatusBadRequest)
            return
        }

        userID, err := strconv.Atoi(r.Header.Get("User-ID"))
        if err != nil {
            http.Error(w, "Invalid user ID", http.StatusBadRequest)
            return
        }

        like := &dbmodel.Like{
            IDPost: postID,
            IDUser: userID,
        }

        if _, err := cfg.LikeRepository.Create(like); err != nil {
            http.Error(w, "Failed to like post", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
    }
}

func UnlikePost(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid post ID", http.StatusBadRequest)
            return
        }

        userID, err := strconv.Atoi(r.Header.Get("User-ID"))
        if err != nil {
            http.Error(w, "Invalid user ID", http.StatusBadRequest)
            return
        }

        if err := cfg.LikeRepository.Delete(postID, userID); err != nil {
            http.Error(w, "Failed to unlike post", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusNoContent)
    }
}