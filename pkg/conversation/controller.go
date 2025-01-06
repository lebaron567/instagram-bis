package conversation

import (
    "encoding/json"
    "net/http"
    "strconv"

    "instagram-bis/config"
    "instagram-bis/database/dbmodel"

    "github.com/go-chi/chi/v5"
)

// CreateDiscussion godoc
// @Summary Create a new discussion
// @Description Create a new discussion
// @Tags discussions
// @Accept json
// @Produce json
// @Param discussion body dbmodel.Discussion true "Discussion"
// @Success 201 {object} dbmodel.Discussion
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to create discussion"
// @Router /discussions [post]
func CreateDiscussion(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var discussion dbmodel.Discussion
        if err := json.NewDecoder(r.Body).Decode(&discussion); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        createdDiscussion, err := cfg.DiscussionRepository.Create(&discussion)
        if err != nil {
            http.Error(w, "Failed to create discussion", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(createdDiscussion)
    }
}

// GetDiscussionsByUserID godoc
// @Summary Get discussions by user ID
// @Description Get discussions by user ID
// @Tags discussions
// @Produce json
// @Param userID path int true "User ID"
// @Success 200 {array} dbmodel.Discussion
// @Failure 400 {string} string "Invalid user ID"
// @Failure 500 {string} string "Failed to retrieve discussions"
// @Router /users/{userID}/discussions [get]
func GetDiscussionsByUserID(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID, err := strconv.Atoi(chi.URLParam(r, "userID"))
        if err != nil {
            http.Error(w, "Invalid user ID", http.StatusBadRequest)
            return
        }

        discussions, err := cfg.DiscussionRepository.FindByUserID(userID)
        if err != nil {
            http.Error(w, "Failed to retrieve discussions", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(discussions)
    }
}

// GetDiscussionByID godoc
// @Summary Get a discussion by ID
// @Description Get a discussion by ID
// @Tags discussions
// @Produce json
// @Param id path int true "Discussion ID"
// @Success 200 {object} dbmodel.Discussion
// @Failure 400 {string} string "Invalid discussion ID"
// @Failure 500 {string} string "Failed to retrieve discussion"
// @Router /discussions/{id} [get]
func GetDiscussionByID(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        discussionID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid discussion ID", http.StatusBadRequest)
            return
        }

        discussion, err := cfg.DiscussionRepository.FindByID(discussionID)
        if err != nil {
            http.Error(w, "Failed to retrieve discussion", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(discussion)
    }
}

// DeleteDiscussion godoc
// @Summary Delete a discussion by ID
// @Description Delete a discussion by ID
// @Tags discussions
// @Param id path int true "Discussion ID"
// @Success 204
// @Failure 400 {string} string "Invalid discussion ID"
// @Failure 500 {string} string "Failed to delete discussion"
// @Router /discussions/{id} [delete]
func DeleteDiscussion(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        discussionID, err := strconv.Atoi(chi.URLParam(r, "id"))
        if err != nil {
            http.Error(w, "Invalid discussion ID", http.StatusBadRequest)
            return
        }

        // Supprimer la discussion
        if err := cfg.DiscussionRepository.Delete(discussionID); err != nil {
            http.Error(w, "Failed to delete discussion", http.StatusInternalServerError)
            return
        }

        // Répondre avec succès
        w.WriteHeader(http.StatusNoContent)
    }
}