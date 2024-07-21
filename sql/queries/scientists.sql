-- name: CreateScientist :one
INSERT INTO cientificos (dni,nomApels)
VALUES ($1, $2)
RETURNING *;