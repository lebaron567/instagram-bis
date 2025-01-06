// filepath: /c:/Users/emeri/Documents/cour ynov/zpi/D-veloppement-d-API/instagram-bis/pkg/user/controller.go
package user

import (
    "encoding/json"
    "net/http"
    "strconv"

    "instagram-bis/config"
    "instagram-bis/database/dbmodel"

    "github.com/go-chi/chi/v5"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param user body dbmodel.User true "User"
// @Success 201 {object} dbmodel.User
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to register user"
// @Router /users/register [post]
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

// LoginUser godoc
// @Summary Log in a user
// @Description Log in a user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /users/login [post]
func LoginUser(cfg *config.Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Implémentation de la connexion et génération de JWT
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]string{"message": "Successfully logged in"})
    }
}

// GetUserProfile godoc
// @Summary Get a user profile by ID
// @Description Get a user profile by ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dbmodel.User
// @Failure 400 {string} string "Invalid user ID"
// @Failure 404 {string} string "User not found"
// @Router /users/{id} [get]
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

// UpdateUserProfile godoc
// @Summary Update a user profile
// @Description Update a user profile
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body dbmodel.User true "User"
// @Success 200 {object} dbmodel.User
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Failed to update user profile"
// @Router /users/{id} [put]
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

// FollowUser godoc
// @Summary Follow a user
// @Description Follow a user
// @Tags users
// @Param id path int true "User ID"
// @Param Current-User-ID header int true "Current User ID"
// @Success 200 {object} map[string]string
// @Failure 400 {string} string "Invalid user ID or current user ID"
// @Failure 500 {string} string "Failed to follow user"
// @Router /users/{id}/follow [post]
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

        // Vérifier si l'utilisateur essaie de se suivre lui-même
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

// GetFollowers godoc
// @Summary Get followers of a user
// @Description Get followers of a user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} dbmodel.Follower
// @Failure 400 {string} string "Invalid user ID"
// @Failure 500 {string} string "Failed to get followers"
// @Router /users/{id}/followers [get]
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

// GetFollowing godoc
// @Summary Get users followed by a user
// @Description Get users followed by a user
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} dbmodel.Follower
// @Failure 400 {string} string "Invalid user ID"
// @Failure 500 {string} string "Failed to get following"
// @Router /users/{id}/following [get]
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