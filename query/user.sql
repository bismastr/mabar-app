-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: InsertUser :exec
INSERT INTO users (
    username,
    avatar_url,
    discord_uid
) VALUES ( $1, $2, $3);