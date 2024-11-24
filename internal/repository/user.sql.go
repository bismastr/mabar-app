// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package repository

import (
	"context"
)

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, avatar_url, discord_uid FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.AvatarUrl,
		&i.DiscordUid,
	)
	return i, err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO users (
    username,
    avatar_url,
    discord_uid
) VALUES ( $1, $2, $3)
`

type InsertUserParams struct {
	Username   string
	AvatarUrl  string
	DiscordUid int64
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.Exec(ctx, insertUser, arg.Username, arg.AvatarUrl, arg.DiscordUid)
	return err
}