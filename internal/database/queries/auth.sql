-- internal/database/queries/auth.sql

-- name: CreateUser :one
INSERT INTO users (email, password_hash, name, lastname)
VALUES ($1, $2, $3, $4)
RETURNING id, email, name, lastname;

-- name: UserExists :one
SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);

-- name: GetUserByEmail :one
SELECT id, email, password_hash, name, lastname
FROM users WHERE email = $1;