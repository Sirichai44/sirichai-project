package drivers

import (
	"context"
	"fmt"
	"log/slog"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	Auth *auth.Client
	DB   *firestore.Client
}

func FirebaseConn(credentialsFilePath string) (*FirebaseClient, error) {
	opt := option.WithCredentialsFile(credentialsFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		slog.Error(fmt.Sprintf("error initializing app: %v\n", err))
		return nil, err
	}
	slog.Info("Firebase initialized")

	client, err := app.Auth(context.Background())
	if err != nil {
		slog.Error(fmt.Sprintf("error getting Auth client: %v\n", err))
		return nil, err
	}
	slog.Info("Firebase Auth initialized")

	firebaseDB, err := app.Firestore(context.Background())
	if err != nil {
		slog.Error(fmt.Sprintf("error getting Database client: %v\n", err))
		return nil, err
	}
	slog.Info("Firebase firestore initialized")

	return &FirebaseClient{
		Auth: client,
		DB:   firebaseDB,
	}, nil
}
