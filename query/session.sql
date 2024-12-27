-- name: InsertGamingSession :one
INSERT INTO sessions (is_finish, session_end, session_start, created_by, game_id) 
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: InsertUserJoinSession :exec
INSERT INTO users_session (user_id, session_id) 
VALUES ($1, $2);

-- name: GetSessionById :many
SELECT 
    s.id AS session_id,
    s.is_finish,
    s.session_end,
    s.session_start,
    s.created_by,
    s.game_id,
    u.id AS user_id,
    u.username,
    u.avatar_url,
    u.discord_uid,
	g.game_name,
	g.game_icon_url,
	c.id AS created_by_user_id,
	c.username AS created_by_username,
	c.avatar_url AS created_by_avatar_url,
	c.discord_uid AS created_by_discord_uid
FROM 
    sessions s
LEFT JOIN
    games g ON s.game_id = g.id
LEFT JOIN
	users c on s.created_by = c.id
LEFT JOIN 
    users_session us ON s.id = us.session_id
LEFT JOIN 
    users u ON us.user_id = u.id
WHERE 
    s.id = $1;

-- name: GetAllSessions :many
SELECT 
    s.id AS session_id,
    s.is_finish,
    s.session_end,
    s.session_start,
    s.created_by,
    s.game_id,
    u.id AS user_id,
    u.username,
    u.avatar_url,
    u.discord_uid,
	g.game_name,
	g.game_icon_url,
	c.id AS created_by_user_id,
	c.username AS created_by_username,
	c.avatar_url AS created_by_avatar_url,
	c.discord_uid AS created_by_discord_uid
FROM 
    sessions s
LEFT JOIN
    games g ON s.game_id = g.id
LEFT JOIN
	users c on s.created_by = c.id
LEFT JOIN 
    users_session us ON s.id = us.session_id
LEFT JOIN 
    users u ON us.user_id = u.id
ORDER BY 
    s.id DESC
LIMIT $1 OFFSET $2;