-- name: GetUser :one
SELECT
	*
FROM
	user
WHERE
	id = ?
LIMIT
	1;

-- name: GetUserByEmail :one
SELECT
	*
FROM
	user
WHERE
	email = ?
LIMIT
	1;

-- name: ListUsers :many
SELECT
	*
FROM
	user
ORDER BY
	name;

-- name: CreateUser :one
INSERT INTO
	user (name, email, password)
VALUES
	(?, ?, ?) RETURNING *;

-- name: UpdateUser :one
UPDATE
	user
set
	name = ?,
	email = ?,
	password = ?
WHERE
	id = ? RETURNING *;

-- name: DeleteUser :exec
DELETE FROM
	user
WHERE
	id = ?;