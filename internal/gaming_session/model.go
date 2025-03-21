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
	ChannelID    string           `json:"channel_id"`
	Name         pgtype.Text      `json:"name"`
}

type CreateGamingSessionResponse struct {
	IsFinish     pgtype.Bool      `json:"is_finish"`
	SessionEnd   pgtype.Timestamp `json:"session_end"`
	SessionStart pgtype.Timestamp `json:"session_start"`
	CreatedBy    int64            `json:"created_by"`
	GameID       int64            `json:"game_id"`
	Name         pgtype.Text      `json:"name"`
	ID           int64            `json:"id"`
}

// JoinGamingSession
type JoinGamingSesionRequest struct {
	UserId    int64 `json:"user_id"`
	SessionId int64 `json:"session_id"`
}

// GetGamingSession
type GetGamingSessionResponse struct {
	SessionID    int64               `json:"session_id"`
	IsFinish     pgtype.Bool         `json:"is_finish"`
	SessionEnd   pgtype.Timestamp    `json:"session_end"`
	SessionStart pgtype.Timestamp    `json:"session_start"`
	Name         pgtype.Text         `json:"name"`
	CreatedBy    GamingSessionUser   `json:"created_by"`
	Users        []GamingSessionUser `json:"members,omitempty"`
	Game         GamingSessionGame   `json:"game_info,omitempty"`
}

type GamingSessionUser struct {
	UserID     pgtype.Int8 `json:"user_id,omitempty"`
	Username   pgtype.Text `json:"username,omitempty"`
	AvatarUrl  pgtype.Text `json:"avatar_url,omitempty"`
	DiscordUid pgtype.Int8 `json:"discord_uid,omitempty"`
}

type GamingSessionGame struct {
	GameId      int64       `json:"id"`
	GameName    pgtype.Text `json:"name"`
	GameIconUrl pgtype.Text `json:"icon_url"`
}

type GetAllGamingSessionRequest struct {
	Page int `json:"page"`
	Rows int `json:"rows"`
}

// Games
type GetAllGamesResponse struct {
	ID          int64  `json:"id"`
	GameName    string `json:"game_name"`
	GameIconUrl string `json:"game_icon_url"`
}
