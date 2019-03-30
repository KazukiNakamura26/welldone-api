package authfirebase

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	initfirebase "github.com/KazukiNakamura26/welldone-api/pkg/init"
	"firebase.google.com/go/auth"
)

type authfirebase struct{}

//User認証情報を取得
func AuthenticationUser() (*auth.Client, error) {
	auth, err := initfirebase.InitFirebase().Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return auth, nil
}

//User認証トークンを取得
func GetAuthUserTokenInfo(w http.ResponseWriter, r *http.Request) (*auth.Token, error) {
	// Firebase SDK のセットアップ
	auth, err := AuthenticationUser()
	if err != nil {
		fmt.Printf("error AuthenticationUser Not: %v\n", err)
		os.Exit(1)
	}

	// クライアントから送られてきた JWT 取得
	authHeader := r.Header.Get("Authorization")
	idToken := strings.Replace(authHeader, "Bearer ", "", 1)
	// JWT の検証
	token, err := auth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Firebase SDK のセットアップ
		token, err := GetAuthUserTokenInfo(w, r)
		if err != nil {
			// JWT が無効なら Handler に進まず別処理
			fmt.Printf("error verifying ID token: %v\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("error verifying ID token\n"))
			return
		}
		_ = token
		next.ServeHTTP(w, r)
	}
}
