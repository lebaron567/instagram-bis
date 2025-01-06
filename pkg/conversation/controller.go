package conversation

import (
	"encoding/json"
	"net/http"
	"strconv"

	"instagram-bis/config"
	"instagram-bis/database/dbmodel"

	"github.com/go-chi/chi/v5"
)

// CreateDiscussion crée une nouvelle discussion
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

// GetDiscussionsByUserID récupère toutes les discussions pour un utilisateur
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

// GetDiscussionByID récupère une discussion par son ID
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

// DeleteDiscussion supprime une discussion par son ID
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
