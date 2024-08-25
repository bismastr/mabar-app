package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/bismastr/discord-bot/model"
	"google.golang.org/api/option"
)

type DbClient struct {
	Client *firestore.Client
}

func NewFirebaseClient(ctx context.Context) *DbClient {
	credsPath := os.Getenv("FIREBASE_CREDENTIALS_PATH")

	creds, err := os.ReadFile(credsPath)
	if err != nil {
		log.Fatalf("Failed to read credentials file: %v", err)
	}

	sa := option.WithCredentialsJSON(creds)
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Firebase Connected")
	return &DbClient{
		Client: client,
	}
}

func (d *DbClient) CreateGamingSession(ctx context.Context, req model.GamingSession) {
	client := d.Client

	_, _, err := client.Collection("gaming-sessions").Add(ctx, req)
	if err != nil {
		log.Printf("An error has occurred when creating gamingsession: %s", err)
	}
}
