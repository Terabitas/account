package domain

import (
	"github.com/nildev/account/Godeps/_workspace/src/golang.org/x/oauth2"
	"github.com/nildev/account/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
)

type (
	// Account is base model
	Account struct {
		Id       string        `bson:"_id" json:"id"`
		Email    string        `bson:"email" json:"email"`
		Username string        `bson:"userName" json:"userName"`
		Avatar   string        `bson:"avatar" json:"avatar"`
		Token    *oauth2.Token `bson:"token" json:"token"`
		Data     Data          `bson:"data" json:"data"`
	}

	// Data used to store arbitrary info of account
	// end users will have different models
	Data map[string]interface{}
)

// MakeAccount constructor
func MakeAccount(email, username, avatar string, token *oauth2.Token, data Data) *Account {
	return &Account{
		Id:       string(bson.NewObjectId()),
		Email:    email,
		Username: username,
		Token:    token,
		Data:     data,
	}
}
