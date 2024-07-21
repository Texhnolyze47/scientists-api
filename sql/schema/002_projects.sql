-- +goose Up
CREATE TABLE proyectos (
    id varchar(10) PRIMARY KEY NOT NULL,
    nombre varchar(100) NOT NULL,
    horas int NOT NULL
);

-- +goose Down
DROP TABLE proyectos;