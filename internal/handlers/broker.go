package handlers

import (
	"credit_broker/internal/services"
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	broker := services.New()
	defer broker.Close()

	if err := broker.Login(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Print(user)
}
