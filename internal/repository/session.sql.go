// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: session.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getAllSession = `-- name: GetAllSession :many
SELECT id, is_finish, session_end, session_start, created_by, game_id FROM sessions 
LIMIT 10
`

func (q *Queries) GetAllSession(ctx context.Context) ([]Session, error) {
	rows, err := q.db.Query(ctx, getAllSession)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Session
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.IsFinish,
			&i.SessionEnd,
			&i.SessionStart,
			&i.CreatedBy,
			&i.GameID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSessionById = `-- name: GetSessionById :one
SELECT id, is_finish, session_end, session_start, created_by, game_id FROM sessions 
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetSessionById(ctx context.Context, id int64) (Session, error) {
	row := q.db.QueryRow(ctx, getSessionById, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.IsFinish,
		&i.SessionEnd,
		&i.SessionStart,
		&i.CreatedBy,
		&i.GameID,
	)
	return i, err
}

const insertGamingSession = `-- name: InsertGamingSession :one
INSERT INTO sessions (is_finish, session_end, session_start, created_by, game_id) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING id, is_finish, session_end, session_start, created_by, game_id
`

type InsertGamingSessionParams struct {
	IsFinish     pgtype.Bool
	SessionEnd   pgtype.Timestamp
	SessionStart pgtype.Timestamp
	CreatedBy    int64
	GameID       int64
}

func (q *Queries) InsertGamingSession(ctx context.Context, arg InsertGamingSessionParams) (Session, error) {
	row := q.db.QueryRow(ctx, insertGamingSession,
		arg.IsFinish,
		arg.SessionEnd,
		arg.SessionStart,
		arg.CreatedBy,
		arg.GameID,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.IsFinish,
		&i.SessionEnd,
		&i.SessionStart,
		&i.CreatedBy,
		&i.GameID,
	)
	return i, err
}

const insertUserJoinSession = `-- name: InsertUserJoinSession :exec
INSERT INTO users_session (user_id, session_id) 
VALUES ($1, $2)
`

type InsertUserJoinSessionParams struct {
	UserID    int64
	SessionID int64
}

func (q *Queries) InsertUserJoinSession(ctx context.Context, arg InsertUserJoinSessionParams) error {
	_, err := q.db.Exec(ctx, insertUserJoinSession, arg.UserID, arg.SessionID)
	return err
}
