package initfirebase

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type initFireBase struct{}

func InitFirebase() *firebase.App {
	// Firebase SDK のセットアップ
	opt := option.WithCredentialsFile("config/welldone-21fc5-firebase-adminsdk-xs05i-63adcdbd60.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("firebase set up error!: %v\n", err)
		os.Exit(1)
	}
	return app
}

//DB接続初期設定(コネクション貼ってる)
func InitFireStore() (*firestore.Client, error) {
	client, err := InitFirebase().Firestore(context.Background())
	if err != nil {
		fmt.Printf("firestore setup error!\n")
		return nil, err
	}
	return client, nil
}
