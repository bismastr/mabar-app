package gaming_session

import (
	"github.com/jackc/pgx/v5/pgtype"
)

// CreateGamingSession
type CreateGamingSessionRequest struct {
	IsFinish     pgtype.Bool      `json:"is_finish"`
	SessionEnd   pgtype.Timestamp `json:"session_end"`
	SessionStart pgtype.Timestamp `json:"session_start"`
	CreatedBy    int64            `json:"created_by"`
	GameID       int64            `json:"game_id"`
}

// JoinGamingSession
type JoinGamingSesionRequest struct {
	UserId    int64 `json:"user_id"`
	SessionId int64 `json:"session_id"`
}

// GetGamingSession
type GetGamingSessionByIdResponse struct {
	SessionID    int64
	IsFinish     pgtype.Bool
	SessionEnd   pgtype.Timestamp
	SessionStart pgtype.Timestamp
	CreatedBy    GamingSessionUser
	GameID       int64
	Users        []GamingSessionUser `json:"users"`
}

type GamingSessionUser struct {
	UserID     pgtype.Int8
	Username   pgtype.Text
	AvatarUrl  pgtype.Text
	DiscordUid pgtype.Int8
}
