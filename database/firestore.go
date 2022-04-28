package database

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/ttvs-blockchain/firebaseupdate/config"
	"google.golang.org/api/option"
)

func InitFireStore(c context.Context, conf *config.Config) (*firestore.Client, error) {
	opt := option.WithCredentialsFile(conf.FirebaseServieAccountKeyPath)
	app, err := firebase.NewApp(c, nil, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(c)
	if err != nil {
		return nil, err
	}
	return client, nil
}
