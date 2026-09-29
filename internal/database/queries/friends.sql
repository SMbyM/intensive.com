-- name: CreateFriendship :exec
INSERT INTO friends (fst, snd) VALUES ($1, $2);

-- name: GetFriendList :many
SELECT snd FROM friends WHERE fst = $1;

-- name: DeleteFriendship :exec
DELETE FROM friends WHERE fst = $1 AND snd = $2;

-- name: AreFriends :one
SELECT EXISTS(SELECT 1 FROM friends WHERE fst = $1 AND snd = $2);