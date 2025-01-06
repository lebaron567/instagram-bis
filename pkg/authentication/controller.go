package authentication

import (
	"encoding/json"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"instagram-bis/database/dbmodel"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		userRepo := dbmodel.NewUserRepository(db)
		passwordHash, err := userRepo.FindPasswordByEmail(req.Email)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		// Générer le token
		user, err := userRepo.FindByEmail(req.Email)
		if err != nil {
			http.Error(w, "User not found", http.StatusInternalServerError)
			return
		}

		token, err := GenerateJWT(int(user.ID), user.Pseudo, user.Email)
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LoginResponse{Token: token})
	}
}
