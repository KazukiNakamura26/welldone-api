package top

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	praiseInfo "github.com/KazukiNakamura26/welldone-api/internal/praise"
	userInfo "github.com/KazukiNakamura26/welldone-api/internal/user"
	authfirebase "github.com/KazukiNakamura26/welldone-api/pkg/auth"
)

type topPageRespons struct {
	UserName     string         `json:"name"`
	ProjectName  map[int]string `json:"projectname"`
	LatestPraise map[int]Praise `json:"latestprice"`
}

type Praise struct {
	Index    int    `json:"id"`
	Praise   string `json:"text"`
	UserName string `json:"praiseName"`
}

func TopPage(w http.ResponseWriter, r *http.Request) {

	// トークン取得
	token, err := authfirebase.GetAuthUserTokenInfo(w, r)
	if err != nil {
		fmt.Printf("private")
		// JWT が無効なら Handler に進まず別処理
		fmt.Printf("error verifying ID token: %v\n", err)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("error verifying ID token\n"))
		return
	}
	// レスポンス変数宣言
	var res topPageRespons
	// 取得したトークンよりユーザー情報を取得
	var user_id = token.UID
	userInfo, err := userInfo.GetUser(user_id)
	if err != nil {
		fmt.Printf("firebase set up error!: %v\n", err)
		os.Exit(1)
	}

	// レスポンスにユーザー名を格納
	res.UserName = userInfo.UserName

	// 褒められ情報取得
	praiseInfo, err := praiseInfo.GetTopPageBasicPraiseInfo(user_id)
	_ = praiseInfo
	_ = err
	// var p Praises
	// for i, v := range praiseInfo {
	// 	var praiseResult Praise
	// 	praiseResult.Index = i
	// 	praiseResult.Praise = v.PraiseText
	// 	praiseResult.UserName = v.PraiseUserName

	// 	p = append(p, doPraises(praiseResult))
	// }
	fmt.Println(res.LatestPraise)

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
