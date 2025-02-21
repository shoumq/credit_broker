package main

import (
	"credit_broker/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/auth/login", withCORS(handlers.Login))
	router.HandleFunc("/api/auth/register", withCORS(handlers.Register))
	router.HandleFunc("/api/auth/is_admin", withCORS(handlers.IsAdmin))

	if err := http.ListenAndServe(":8085", router); err != nil {
		panic(err)
	}
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
