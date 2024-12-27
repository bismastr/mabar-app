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
		Name:         gamingSession.Name,
	})

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (g *GamingSessionService) GetGamingSessionById(ctx context.Context, id int64) (*GetGamingSessionResponse, error) {
	result, err := g.repository.GetSessionById(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(result) < 1 {
		return nil, errors.New("cannot find session")
	}

	response := GetGamingSessionResponse{
		SessionID:    result[0].SessionID,
		IsFinish:     result[0].IsFinish,
		SessionEnd:   result[0].SessionEnd,
		SessionStart: result[0].SessionStart,
		CreatedBy: GamingSessionUser{
			UserID:     result[0].CreatedByUserID,
			DiscordUid: result[0].CreatedByDiscordUid,
			AvatarUrl:  result[0].CreatedByAvatarUrl,
			Username:   result[0].CreatedByUsername,
		},
		Game: GamingSessionGame{
			GameId:      result[0].GameID,
			GameName:    result[0].GameName,
			GameIconUrl: result[0].GameIconUrl,
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

func (g *GamingSessionService) GetAllGamingSessions(ctx context.Context, req *GetAllGamingSessionRequest) (*[]GetGamingSessionResponse, error) {
	offset := (req.Page - 1) * req.Rows
	rows, err := g.repository.GetAllSessions(ctx, repository.GetAllSessionsParams{
		Limit:  int32(req.Rows),
		Offset: int32(offset),
	})

	if err != nil {
		return nil, err
	}

	var response []GetGamingSessionResponse
	for _, row := range rows {
		if len(response) < 1 || row.SessionID != response[len(response)-1].SessionID {
			response = append(response, GetGamingSessionResponse{
				SessionID:    row.SessionID,
				IsFinish:     row.IsFinish,
				SessionEnd:   row.SessionEnd,
				SessionStart: row.SessionStart,
				Name:         row.Name,
				CreatedBy: GamingSessionUser{
					UserID:     row.CreatedByUserID,
					DiscordUid: row.CreatedByDiscordUid,
					AvatarUrl:  row.CreatedByAvatarUrl,
					Username:   row.CreatedByUsername,
				},
				Game: GamingSessionGame{
					GameId:      row.GameID,
					GameName:    row.GameName,
					GameIconUrl: row.GameIconUrl,
				},
			})

		}

		if row.UserID.Valid {
			response[len(response)-1].Users = append(response[len(response)-1].Users, GamingSessionUser{
				UserID:     row.UserID,
				DiscordUid: row.DiscordUid,
				AvatarUrl:  row.AvatarUrl,
				Username:   row.Username,
			})
		}
	}

	return &response, nil
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

func (g *GamingSessionService) GetAllGames(ctx context.Context) (*[]repository.Game, error) {
	row, err := g.repository.GetGameList(ctx)
	if err != nil {
		return nil, err
	}

	return &row, nil
}
