package users

import (
	"context"
	"fmt"

	user "github.com/KazukiNakamura26/welldone-api/internal/user"
	initFireBase "github.com/KazukiNakamura26/welldone-api/pkg/init"
	"google.golang.org/api/iterator"
)

// ユーザーID1つに対しての情報を返却
type BasicPraiseInfoData struct {
	PraiseText   string `firestore:"praise_txt"`
	PraiseUserId string `firestore:"praise_user_id"`
	ProjectId    string `firestore:"project_id"`
	UserId       string `firestore:"user_id"`
}

type ReturnTopPagePraise struct {
	PraiseText     string
	PraiseUserName string
}

type ReturnTopPagePraises []*ReturnTopPagePraise

func doReturnTopPagePraise(s ReturnTopPagePraise) (r *ReturnTopPagePraise) {
	r = new(ReturnTopPagePraise)
	r.PraiseText = s.PraiseText
	r.PraiseUserName = s.PraiseUserName

	return r
}

// ユーザーID1つに対して褒められた情報を返却
func GetTopPageBasicPraiseInfo(userId string) (ReturnTopPagePraises, error) {
	// 返却用の変数宣言
	var topPages ReturnTopPagePraises
	//　firestoer初期設定
	client, err := initFireBase.InitFireStore()
	if err != nil {
		fmt.Printf("firestore client error!\n")
		// return nil, err
	}
	clients := client.Collection("praise")

	info := clients.Where("user_id", "==", userId).Documents(context.Background())
	//map = info.GetAll()

	for {
		var result ReturnTopPagePraise
		doc, err := info.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Printf("document error!\n : %v\n", err)
		}

		var info BasicPraiseInfoData
		doc.DataTo(&info)
		//fmt.Println(info.PraiseText)
		// 取得した値をstringに変換
		result.PraiseText = info.PraiseText

		//褒めたユーザー名取得
		uid := info.PraiseUserId
		userInfo, err := user.GetUser(uid)
		result.PraiseUserName = userInfo.UserName

		// 返却用の変数に格納
		topPages = append(topPages, doReturnTopPagePraise(result))
	}

	return topPages, nil

}
