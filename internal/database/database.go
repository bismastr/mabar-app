package database

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type DbClient struct {
	Client *firestore.Client
}

func NewFirebaseClient(ctx context.Context) (*DbClient, error) {
	credentialsBase64 := os.Getenv("FIREBASE_CREDENTIALS")

	credentials, err := base64.StdEncoding.DecodeString(credentialsBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %w", err)
	}

	firestoreApp, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(credentials))
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore app: %w", err)
	}

	firestoreClient, err := firestoreApp.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore client: %w", err)
	}

	return &DbClient{
		Client: firestoreClient,
	}, nil
}
