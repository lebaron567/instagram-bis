package user

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/go-chi/chi/v5"
    "instagram-bis/config"
    "instagram-bis/database/dbmodel"
)

func RegisterUser(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user dbmodel.User
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }

        if _, err := cfg.UserRepository.Create(&user); err != nil {
            http.Error(w, "Failed to register user", http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(user)
    }
}

func LoginUser(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Implémentation de la connexion et génération de JWT
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
        // Implémentation pour suivre un autre utilisateur
    }
}

func GetFollowers(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Implémentation pour obtenir la liste des abonnés
    }
}

func GetFollowing(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Implémentation pour obtenir la liste des abonnements
    }
}