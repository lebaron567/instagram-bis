// filepath: /c:/Users/emeri/Documents/cour ynov/zpi/D-veloppement-d-API/instagram-bis/pkg/user/controller.go
package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"instagram-bis/config"
	"instagram-bis/database/dbmodel"

	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/chi/v5"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}


func RegisterUser(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user dbmodel.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword

		if _, err := cfg.UserRepository.Create(&user); err != nil {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			return
		}

		reponse := dbmodel.User{
			LastName:  user.LastName,
			FirstName: user.FirstName,
			Email:     user.Email,
			Pseudo:    user.Pseudo,
			Birthdate: user.Birthdate,
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(reponse)
	}
}


func LoginUser(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Successfully logged in"})
	}
}


func GetUserProfile(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		user, err := cfg.UserRepository.FindByID(userID)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}


func UpdateUserProfile(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var updatedUser dbmodel.User
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		user, err := cfg.UserRepository.UpdateUser(userID, &updatedUser)
		if err != nil {
			http.Error(w, "Failed to update user profile", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

func FollowUser(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Récupérer l'ID de l'utilisateur à suivre
		userID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Récupérer l'ID de l'utilisateur actuel (par exemple, à partir du contexte ou du token JWT)
		currentUserID, err := strconv.Atoi(r.Header.Get("Current-User-ID"))
		if err != nil {
			http.Error(w, "Invalid current user ID", http.StatusBadRequest)
			return
		}

		if userID == currentUserID {
			http.Error(w, "You cannot follow yourself", http.StatusBadRequest)
			return
		}

		// Créer une nouvelle relation de suivi
		follow := &dbmodel.Follower{
			IDUser:     userID,
			IDFollower: currentUserID,
		}

		// Enregistrer la relation de suivi dans la base de données
		if _, err := cfg.FollowerRepository.Follow(follow); err != nil {
			http.Error(w, "Failed to follow user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Successfully followed user"})
	}
}


func GetFollowers(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implémentation pour obtenir la liste des abonnés
		userID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		followers, err := cfg.FollowerRepository.FindFollowersByUserID(userID)
		if err != nil {
			http.Error(w, "Failed to get followers", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(followers)
	}
}


func GetFollowing(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implémentation pour obtenir la liste des abonnements
		userID, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		following, err := cfg.FollowerRepository.FindFollowingByUserID(userID)
		if err != nil {
			http.Error(w, "Failed to get following", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(following)
	}
}
