-- name: CreateProject :one
INSERT INTO proyectos (id,nombre,horas)
VALUES ($1, $2, $3)
RETURNING *;