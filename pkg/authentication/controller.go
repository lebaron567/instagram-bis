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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if req.Email != mockUser.Email || req.Password != mockUser.Password {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(mockUser.ID, mockUser.Pseudo, mockUser.Email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: token})
}
