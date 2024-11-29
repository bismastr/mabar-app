-- name: GetGameList :many
SELECT games.id, games.game_name, games.game_icon_url FROM games
LIMIT 10;