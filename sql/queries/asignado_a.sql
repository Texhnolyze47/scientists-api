-- name: CreateProyectoAsignado :one
INSERT INTO asignado_a (cientifico,proyecto,dedicacion)
VALUES ($1, $2, $3)
RETURNING *;

