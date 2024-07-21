-- +goose Up
UPDATE asignado_a
SET dedicacion = 0;
-- +goose Down
