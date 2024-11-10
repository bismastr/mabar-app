package gamingSession

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/bismastr/discord-bot/internal/database"
	"google.golang.org/api/iterator"
)

type FirestoreRepositorySession interface {
	CreateGamingSession(ctx context.Context, req GamingSession) (string, error)
	GetAllGamingSession(ctx context.Context) (*[]GamingSession, error)
	GetGamingSessionByRefId(ctx context.Context, id string) (*GamingSession, error)
	UpdateGamingSessionByRefId(ctx context.Context, id string, req GamingSession) error
}

type FirestoreRepositorySessionImpl struct {
	DbClient *database.DbClient
}

func NewRepositoryImpl(dbClient *database.DbClient) *FirestoreRepositorySessionImpl {
	return &FirestoreRepositorySessionImpl{
		DbClient: dbClient,
	}
}

func (r *FirestoreRepositorySessionImpl) CreateGamingSession(ctx context.Context, req GamingSession) (string, error) {
	docRef, _, err := r.DbClient.Client.Collection("gaming-sessions").Add(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to create gaming session: %w", err)
	}

	return docRef.ID, nil
}

func (r *FirestoreRepositorySessionImpl) GetAllGamingSession(ctx context.Context) (*[]GamingSession, error) {
	iter := r.DbClient.Client.Collection("gaming-sessions").Documents(ctx)
	defer iter.Stop()

	var result []GamingSession
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}

		if err != nil {
			return nil, err
		}

		var docIteration GamingSession
		doc.DataTo(&docIteration)

		result = append(result, docIteration)

	}

	return &result, nil
}

func (r *FirestoreRepositorySessionImpl) GetGamingSessionByRefId(ctx context.Context, refId string) (*GamingSession, error) {

	docRef := r.DbClient.Client.Collection("gaming-sessions").Doc(refId)

	doc, err := docRef.Get(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get document: %w", err)
	}

	var result GamingSession
	err = doc.DataTo(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to get parse document: %w", err)
	}

	return &result, nil
}

func (r *FirestoreRepositorySessionImpl) UpdateGamingSessionByRefId(ctx context.Context, refId string, req GamingSession) error {
	docRef := r.DbClient.Client.Collection("gaming-sessions").Doc(refId)
	updateValue := []firestore.Update{}
	addUpdate(&updateValue, "game_name", req.GameName)
	addUpdate(&updateValue, "session_start", req.SessionStart)
	addUpdate(&updateValue, "session_end", req.SessionEnd)
	addUpdate(&updateValue, "is_finish", req.IsFinish)
	addUpdate(&updateValue, "created_by", req.CreatedBy)
	addUpdate(&updateValue, "members_sessions", req.MembersSession)
	addUpdate(&updateValue, "created_at", req.CreatedAt)

	_, err := docRef.Update(ctx, updateValue)
	if err != nil {
		return fmt.Errorf("failed to update document: %w", err)
	}

	return err
}
