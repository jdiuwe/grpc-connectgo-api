-- name: CreateUser :one
INSERT INTO users (
    uuid, firstname, lastname, email, hashed_password
) VALUES (
             $1, $2, $3, $4, $5
         )
RETURNING *;

-- name: GetUserByUUID :one
SELECT * FROM users
WHERE uuid = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UserEmailExists :one
SELECT 1 FROM users
WHERE email = $1 LIMIT 1;

-- name: SetUserVerified :exec
UPDATE users
SET verified=true
WHERE id = $1;