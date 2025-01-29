package handlers

import (
	"credit_broker/internal/services"
	"encoding/json"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type UserID struct {
	UserID int64 `json:"user-id"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	broker := services.New()
	defer broker.Close()

	var tokens AuthResponse
	token, err := broker.Login(services.AuthRequest(user))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tokens.Token = token

	if err := json.NewEncoder(w).Encode(tokens); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	broker := services.New()
	defer broker.Close()

	var tokens UserID
	userId, err := broker.Register(services.AuthRequest(user))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tokens.UserID = userId

	if err := json.NewEncoder(w).Encode(tokens); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func IsAdmin(w http.ResponseWriter, r *http.Request) {
	var user AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	broker := services.New()
	defer broker.Close()

	var tokens AuthResponse
	token, err := broker.Login(services.AuthRequest(user))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	tokens.Token = token

	if err := json.NewEncoder(w).Encode(tokens); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
