package main

import (
	"fmt"
	"log"
	"net/http"

	top "../internal/top"
	authfirebase "../pkg/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public!\n"))
	fmt.Printf("public in!!!!!!!!!")
}

func main() {
	port := "3030"
	fmt.Printf("port: " + port)
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

	r := mux.NewRouter()
	r.HandleFunc("/", public)
	r.HandleFunc("/private", authfirebase.AuthMiddleware(top.TopPage))

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}