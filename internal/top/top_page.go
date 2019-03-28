package top

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	userInfo "github.com/KazukiNakamura26/welldone-api/internal/user"
	authfirebase "github.com/KazukiNakamura26/welldone-api/pkg/auth"
)

type topPageRespons struct {
	UserName string `json:"name"`
}

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
	var res topPageRespons
	res.UserName = userInfo.UserName

	topPageResponsHandler(w, r, res)
}

func topPageResponsHandler(w http.ResponseWriter, r *http.Request, topRespons topPageRespons) {
	res, err := json.Marshal(topRespons)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
