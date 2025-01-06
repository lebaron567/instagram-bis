package messagerie

import (
	"encoding/json"
	"net/http"
	"strconv"

	"instagram-bis/config"
	"instagram-bis/database/dbmodel"

	"github.com/go-chi/chi/v5"
)

// CreateMessage crée un nouveau message
func CreateMessage(cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var msg dbmodel.Message
		if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if _, err := cfg.MessageRepository.Create(&msg); err != nil {
			http.Error(w, "Failed to create message", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(msg)
	}
}

// GetMessagesByDiscussion récupère les messages liés à une discussion
func GetMessagesByDiscussion(cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		discussionID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid discussion ID", http.StatusBadRequest)
			return
		}

		messages, err := cfg.MessageRepository.FindByDiscussionID(discussionID)
		if err != nil {
			http.Error(w, "Failed to retrieve messages", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(messages)
	}
}

// UpdateMessage met à jour un message existant
func UpdateMessage(cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		messageID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid message ID", http.StatusBadRequest)
			return
		}

		var updatedMessage dbmodel.Message
		if err := json.NewDecoder(r.Body).Decode(&updatedMessage); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		updatedMsg, err := cfg.MessageRepository.Update(messageID, &updatedMessage) // Utilisation de la méthode Update avec l'ID
		if err != nil {
			http.Error(w, "Failed to update message", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedMsg) // Retourner le message mis à jour
	}
}

// DeleteMessage supprime un message
func DeleteMessage(cfg *config.Config) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		messageID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid message ID", http.StatusBadRequest)
			return
		}

		if err := cfg.MessageRepository.Delete(messageID); err != nil {
			http.Error(w, "Failed to delete message", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
