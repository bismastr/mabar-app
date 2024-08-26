package db

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/bismastr/discord-bot/model"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type DbClient struct {
	Client *firestore.Client
}

func NewFirebaseClient(ctx context.Context) *DbClient {
	creds := os.Getenv("FIREBASE_CREDENTIALS")

	decoded, err := base64.StdEncoding.DecodeString(creds)
	if err != nil {
		log.Fatalf("Failed to decode Base64: %v", err)
	}

	sa := option.WithCredentialsJSON(decoded)
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

func (d *DbClient) CreateGamingSession(ctx context.Context, req model.GamingSession) error {
	client := d.Client

	_, _, err := client.Collection("gaming-sessions").Add(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (d *DbClient) ReadGamingSessionByCreatedUserId(ctx context.Context, id string) (*model.GamingSession, error) {
	client := d.Client

	query := client.Collection("gaming-sessions").Where("created_by.id", "==", id).Limit(1)

	iter := query.Documents(ctx)
	defer iter.Stop()

	doc, err := iter.Next()
	if err == iterator.Done {
		return nil, nil //No session found
	}
	if err != nil {
		return nil, fmt.Errorf("failed to iterate documents: %w", err)
	}

	var session model.GamingSession
	err = doc.DataTo(&session)
	if err != nil {
		return nil, fmt.Errorf("failed to map document data: %w", err)
	}

	return &session, nil
}
