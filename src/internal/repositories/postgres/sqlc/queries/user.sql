-- name: CreateUser :exec

INSERT INTO "user" ("id", "name", "email", "preferences") VALUES ($1, $2, $3, $4);

-- name: ListUsers :many

SELECT * FROM "user";

