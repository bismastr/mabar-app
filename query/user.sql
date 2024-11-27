-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByDiscordUID :one
SELECT * FROM users
WHERE discord_uid = $1 LIMIT 1;

-- name: InsertUser :exec
INSERT INTO users (
    username,
    avatar_url,
    discord_uid
) VALUES ( $1, $2, $3)
ON CONFLICT (discord_uid) DO UPDATE SET
username = EXCLUDED.username,
avatar_url = EXCLUDED.avatar_url;