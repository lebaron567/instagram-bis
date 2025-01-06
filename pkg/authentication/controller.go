package authentication

import (
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Simule un utilisateur pour l'exemple
var mockUser = struct {
	ID       int
	Email    string
	Password string
	Pseudo   string
}{
	ID:       1,
	Email:    "user@example.com",
	Password: "password123",
	Pseudo:   "mockUser",
}

// LoginHandler gère l'authentification
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Vérification des informations d'identification
	if req.Email != mockUser.Email || req.Password != mockUser.Password {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Générer un token JWT
	token, err := GenerateJWT(mockUser.ID, mockUser.Pseudo, mockUser.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Réponse avec le token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
