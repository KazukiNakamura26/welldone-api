package userinfo

import (
	"context"
	"fmt"

	initFireBase "../../pkg/init"
)

type UserInfoDate struct {
	UserName string `firestore:"name"`
}

func GetUser(userId string) (*UserInfoDate, error) {
	//firestoer
	client, err := initFireBase.InitFireStore()
	if err != nil {
		fmt.Printf("firestore client error!\n")
		return nil, err
	}

	info, err := client.Collection("users").Doc(userId).Get(context.Background())
	if err != nil {
		fmt.Printf("firestore user get error!\n")
		return nil, err
	}

	var userInfo UserInfoDate
	info.DataTo(&userInfo)
	return &userInfo, nil
}
