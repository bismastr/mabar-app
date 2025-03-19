-- name: GetItemsContainsName :many
SELECT id, name, hash_name FROM items
WHERE hash_name ILIKE '%' || $1 || '%'
LIMIT 25;