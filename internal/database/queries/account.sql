-- name: GetUserById :one
SELECT id, name, lastname, nickname, email, phone, male, birthday
FROM users
WHERE id = $1;

-- name: UpdateUserProfile :exec
UPDATE users
SET
    name     = COALESCE(sqlc.narg('name'), name),
    lastname = COALESCE(sqlc.narg('lastname'), lastname),
    nickname = COALESCE(sqlc.narg('nickname'), nickname),
    phone    = COALESCE(sqlc.narg('phone'), phone),
    male     = COALESCE(sqlc.narg('male'), male),
    birthday = COALESCE(sqlc.narg('birthday'), birthday)
WHERE id = sqlc.arg('id');