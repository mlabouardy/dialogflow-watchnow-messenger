package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/webhook", VerificationEndPoint).Methods("GET")
	r.HandleFunc("/webhook", MessagesEndPoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
