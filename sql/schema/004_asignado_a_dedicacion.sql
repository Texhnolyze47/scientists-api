-- +goose Up
ALTER TABLE asignado_a
    ADD dedicacion INT;

-- +goose Down
ALTER TABLE asignado_a
    DROP COLUMN dedicacion;