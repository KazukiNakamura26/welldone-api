package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	top "github.com/KazukiNakamura26/welldone-api/internal/top"
	authfirebase "github.com/KazukiNakamura26/welldone-api/pkg/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public!\n"))
	fmt.Printf("public in!!!!!!!!!")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	allowedOrigins := handlers.AllowedOrigins([]string{"https://welldone-app.herokuapp.com"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization"})

	r := mux.NewRouter()
	r.HandleFunc("/", public)
	r.HandleFunc("/private", authfirebase.AuthMiddleware(top.TopPage))

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
