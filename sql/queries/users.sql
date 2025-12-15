-- name: CreateUser :one
INSERT INTO "user"(email,username,password)
VALUES ($1,$2,$3)
RETURNING *;