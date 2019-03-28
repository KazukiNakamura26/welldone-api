package top

import (
	"fmt"
	"net/http"
	"os"

	authfirebase "../../pkg/auth"
	userInfo "../user"
)

func TopPage(w http.ResponseWriter, r *http.Request) {

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

	userInfo, err := userInfo.GetUser(user_id)
	if err != nil {
		fmt.Printf("firebase set up error!: %v\n", err)
		os.Exit(1)
	}

	w.Write([]byte(userInfo.UserName))
}
