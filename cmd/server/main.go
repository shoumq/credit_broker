package main

import (
	"credit_broker/internal/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.Login)
	router.HandleFunc("/register", handlers.Register)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
