package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	userinfo "../internal/user"
	authfirebase "../pkg/auth"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func public(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello public!\n"))
	fmt.Printf("public in!!!!!!!!!")
}

func private(w http.ResponseWriter, r *http.Request) {

	token, err := authfirebase.GetAuthUserTokenInfo(w, r)
	if err != nil {
		fmt.Printf("private")
		// JWT が無効なら Handler に進まず別処理
		fmt.Printf("error verifying ID token: %v\n", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("error verifying ID token\n"))
		return
	}
	var user_id = token.UID

	userInfo, err := userinfo.GetUser(user_id)
	if err != nil {
		fmt.Printf("firebase set up error!: %v\n", err)
		os.Exit(1)
	}

	w.Write([]byte(userInfo.UserName))
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
	r.HandleFunc("/private", authfirebase.AuthMiddleware(private))

	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
