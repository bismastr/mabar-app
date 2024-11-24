package gaming_session

import (
	"context"
	"errors"

	"github.com/bismastr/discord-bot/internal/repository"
)

type GamingSessionService struct {
	repository *repository.Queries
}

func NewGamingSessionService(repository *repository.Queries) *GamingSessionService {
	return &GamingSessionService{
		repository: repository,
	}
}

func (g *GamingSessionService) CreateGamingSession(ctx context.Context, gamingSession *CreateGamingSessionRequest) (*repository.Session, error) {
	result, err := g.repository.InsertGamingSession(ctx, repository.InsertGamingSessionParams{
		IsFinish:     gamingSession.IsFinish,
		SessionEnd:   gamingSession.SessionEnd,
		SessionStart: gamingSession.SessionStart,
		CreatedBy:    gamingSession.CreatedBy,
		GameID:       gamingSession.GameID,
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (g *GamingSessionService) GetGamingSessionById(ctx context.Context, id int64) (*GetGamingSessionByIdResponse, error) {
	result, err := g.repository.GetSessionById(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(result) < 1 {
		return nil, errors.New("cannot find session")
	}

	response := GetGamingSessionByIdResponse{
		SessionID:    result[0].SessionID,
		IsFinish:     result[0].IsFinish,
		SessionEnd:   result[0].SessionEnd,
		SessionStart: result[0].SessionStart,
		GameID:       result[0].GameID,
		CreatedBy: GamingSessionUser{
			UserID:     result[0].CreatedByUserID,
			DiscordUid: result[0].CreatedByDiscordUid,
			AvatarUrl:  result[0].CreatedByAvatarUrl,
			Username:   result[0].CreatedByUsername,
		},
	}

	for _, row := range result {
		response.Users = append(response.Users, GamingSessionUser{
			UserID:     row.UserID,
			DiscordUid: row.DiscordUid,
			AvatarUrl:  row.AvatarUrl,
			Username:   row.Username,
		})
	}

	return &response, nil
}

func (g *GamingSessionService) GetAllGamingSessions(ctx context.Context) (*[]repository.Session, error) {
	result, err := g.repository.GetAllSession(ctx)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (g *GamingSessionService) InsertUserJoinSession(ctx context.Context, userId int64, sessionId int64) error {
	err := g.repository.InsertUserJoinSession(ctx, repository.InsertUserJoinSessionParams{
		UserID:    userId,
		SessionID: sessionId,
	})

	if err != nil {
		return err
	}

	return nil
}
