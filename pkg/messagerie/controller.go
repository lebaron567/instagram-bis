// filepath: /c:/Users/emeri/Documents/cour ynov/zpi/D-veloppement-d-API/instagram-bis/pkg/messagerie/controller.go
package messagerie

import (
	"encoding/json"
	"net/http"
	"strconv"

	"instagram-bis/config"
	"instagram-bis/database/dbmodel"
	model "instagram-bis/pkg/models"

	"github.com/go-chi/chi/v5"
)


func CreateMessage(cfg *config.Config) http.HandlerFunc {
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

		response := model.Message{
			ID:           msg.ID,
			IDDiscussion: msg.IDDiscussion,
			IDUser:       msg.IDUser,
			Content:      msg.Content,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}


func GetMessagesByDiscussion(cfg *config.Config) http.HandlerFunc {
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


func UpdateMessage(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		messageID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid message ID", http.StatusBadRequest)
			return
		}

		var updatedMsg dbmodel.Message
		if err := json.NewDecoder(r.Body).Decode(&updatedMsg); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		updatedMsgPtr, err := cfg.MessageRepository.Update(messageID, &updatedMsg)
		if err != nil {
			http.Error(w, "Failed to update message", http.StatusInternalServerError)
			return
		}
		updatedMsg = *updatedMsgPtr

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedMsg)
	}
}


func DeleteMessage(cfg *config.Config) http.HandlerFunc {
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
