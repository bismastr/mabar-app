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

func (d *DbClient) CreateGamingSession(ctx context.Context, req model.GamingSession) (string, error) {
	client := d.Client

	docRef, _, err := client.Collection("gaming-sessions").Add(ctx, req)
	if err != nil {
		return "", err
	}

	return docRef.ID, nil
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

func (d *DbClient) AddMemberToSession(ctx context.Context, refId string, newMember string) (*model.GamingSession, error) {
	client := d.Client

	// Get the specific document by ID
	docRef := client.Collection("gaming-sessions").Doc(refId)
	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get document: %w", err)
	}

	var session model.GamingSession
	err = doc.DataTo(&session)
	if err != nil {
		return nil, fmt.Errorf("failed to map document data: %w", err)
	}

	// Update member session
	session.MembersSession = append(session.MembersSession, newMember)
	_, err = docRef.Set(ctx, map[string]interface{}{
		"members_sessions": session.MembersSession,
	}, firestore.MergeAll)
	if err != nil {
		return nil, fmt.Errorf("failed to update document: %w", err)
	}

	return &session, nil
}

func (d *DbClient) ReadGamingSession(ctx context.Context) ([]model.GamingSession, error) {
	client := d.Client
	var result []model.GamingSession

	iter := client.Collection("gaming-sessions").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return result, err
		}

		var docIteration model.GamingSession
		doc.DataTo(&docIteration)

		result = append(result, docIteration)
	}

	return result, nil
}

func (d *DbClient) GetMembersList(ctx context.Context, refId string) (*model.GamingSession, error) {
	client := d.Client

	docRef := client.Collection("gaming-sessions").Doc(refId)
	doc, _ := docRef.Get(ctx)

	var m *model.GamingSession
	err := doc.DataTo(&m)
	if err != nil {
		return nil, err
	}

	return m, err
}
