// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: price.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getItemsContainsName = `-- name: GetItemsContainsName :many
SELECT id, name, hash_name FROM items
WHERE hash_name ILIKE '%' || $1 || '%'
LIMIT 25
`

type GetItemsContainsNameRow struct {
	ID       int32
	Name     string
	HashName string
}

func (q *Queries) GetItemsContainsName(ctx context.Context, dollar_1 pgtype.Text) ([]GetItemsContainsNameRow, error) {
	rows, err := q.db.Query(ctx, getItemsContainsName, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetItemsContainsNameRow
	for rows.Next() {
		var i GetItemsContainsNameRow
		if err := rows.Scan(&i.ID, &i.Name, &i.HashName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
