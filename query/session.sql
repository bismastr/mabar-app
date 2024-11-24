-- name: GetAllSession :many
SELECT * FROM sessions 
LIMIT 10;

-- name: InsertGamingSession :one
INSERT INTO sessions (is_finish, session_end, session_start, created_by, game_id) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: InsertUserJoinSession :exec
INSERT INTO users_session (user_id, session_id) 
VALUES ($1, $2);

-- name: GetSessionById :one
SELECT * FROM sessions 
WHERE id = $1
LIMIT 1;