-- name: GetTodo :one
SELECT
	*
FROM
	todo
WHERE
	id = ?
LIMIT
	1;

-- name: ListTodosByOwner :many
SELECT
	*
FROM
	todo
WHERE
	owner_id = ?
ORDER BY
	title;

-- name: CreateTodo :one
INSERT INTO
	todo (title, content, owner_id)
VALUES
	(?, ?, ?) RETURNING *;

-- name: UpdateTodo :one
UPDATE
	todo
set
	title = ?,
	content = ?,
	is_done = ?
WHERE
	id = ? RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM
	todo
WHERE
	id = ?;