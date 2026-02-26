-- name: CreateUser :exec
INSERT INTO users (name, dob)
VALUES (?, ?);

-- name: CreateUserWithAuth :exec
INSERT INTO users (name, email, password_hash, role, dob)
VALUES (?, ?, ?, ?, ?);

-- name: GetUser :one
SELECT id, name, dob FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT id, name, email, password_hash, role, dob, created_at, updated_at FROM users WHERE email = ?;

-- name: GetUserByID :one
SELECT id, name, email, password_hash, role, dob, created_at, updated_at FROM users WHERE id = ?;

-- name: ListUsers :many
SELECT id, name, dob FROM users ORDER BY id;

-- name: UpdateUser :exec
UPDATE users
SET name = ?, dob = ?
WHERE id = ?;

-- name: UpdateUserWithEmail :exec
UPDATE users
SET name = ?, email = ?, dob = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;