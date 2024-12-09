package firebase

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FirebaseClient struct {
	Messaging *messaging.Client
}

func NewFirebaseClient(ctx context.Context) (*FirebaseClient, error) {
	credentialsBase64 := os.Getenv("FIREBASE_CREDENTIALS")

	credentials, err := base64.StdEncoding.DecodeString(credentialsBase64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %w", err)
	}

	firbaseApp, err := firebase.NewApp(ctx, nil, option.WithCredentialsJSON(credentials))
	if err != nil {
		return nil, fmt.Errorf("failed to create firestore app: %w", err)
	}

	messaging, err := firbaseApp.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create firebase messaging: %w", err)
	}

	return &FirebaseClient{
		Messaging: messaging,
	}, nil
}
