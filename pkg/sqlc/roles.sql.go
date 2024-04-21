// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: roles.sql

package sqlc

import (
	"context"
	"database/sql"
	"strings"
)

const createRole = `-- name: CreateRole :execlastid
INSERT INTO roles(created_at, name, description)
VALUES(CURRENT_TIMESTAMP, ?, ?)
`

type CreateRoleParams struct {
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
}

func (q *Queries) CreateRole(ctx context.Context, arg CreateRoleParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, createRole, arg.Name, arg.Description)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const findRoleByName = `-- name: FindRoleByName :one
SELECT id, created_at, updated_at, deleted_at, name, description
FROM roles
where lower(name) = lower(?)
LIMIT 1
`

func (q *Queries) FindRoleByName(ctx context.Context, lower string) (Role, error) {
	row := q.db.QueryRowContext(ctx, findRoleByName, lower)
	var i Role
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
		&i.Name,
		&i.Description,
	)
	return i, err
}

const findRolesByNames = `-- name: FindRolesByNames :many
SELECT id, created_at, updated_at, deleted_at, name, description
FROM roles
where name in (/*SLICE:names*/?)
`

func (q *Queries) FindRolesByNames(ctx context.Context, names []string) ([]Role, error) {
	query := findRolesByNames
	var queryParams []interface{}
	if len(names) > 0 {
		for _, v := range names {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:names*/?", strings.Repeat(",?", len(names))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:names*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.Name,
			&i.Description,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}