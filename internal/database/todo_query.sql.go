// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: todo_query.sql

package database

import (
	"context"
	"database/sql"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO
	todo (title, content, owner_id)
VALUES
	(?, ?, ?) RETURNING id, title, content, is_done, owner_id, created_at, updated_at
`

type CreateTodoParams struct {
	Title   string
	Content sql.NullString
	OwnerID int64
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, createTodo, arg.Title, arg.Content, arg.OwnerID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.IsDone,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM
	todo
WHERE
	id = ?
`

func (q *Queries) DeleteTodo(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTodo, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT
	id, title, content, is_done, owner_id, created_at, updated_at
FROM
	todo
WHERE
	id = ?
LIMIT
	1
`

func (q *Queries) GetTodo(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getTodo, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.IsDone,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listTodosByOwner = `-- name: ListTodosByOwner :many
SELECT
	id, title, content, is_done, owner_id, created_at, updated_at
FROM
	todo
WHERE
	owner_id = ?
ORDER BY
	title
`

func (q *Queries) ListTodosByOwner(ctx context.Context, ownerID int64) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, listTodosByOwner, ownerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.IsDone,
			&i.OwnerID,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateTodo = `-- name: UpdateTodo :one
UPDATE
	todo
set
	title = ?,
	content = ?,
	is_done = ?
WHERE
	id = ? RETURNING id, title, content, is_done, owner_id, created_at, updated_at
`

type UpdateTodoParams struct {
	Title   string
	Content sql.NullString
	IsDone  bool
	ID      int64
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateTodo,
		arg.Title,
		arg.Content,
		arg.IsDone,
		arg.ID,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.IsDone,
		&i.OwnerID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
