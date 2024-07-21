-- +goose Up
CREATE TABLE asignado_a (
    cientifico varchar(10) NOT NULL,
    proyecto varchar(10) NOT NULL,
    PRIMARY KEY (cientifico, proyecto),
    FOREIGN KEY (cientifico) REFERENCES cientificos(dni),
    FOREIGN KEY (proyecto) REFERENCES proyectos(id)
);

-- +goose Down
DROP TABLE asignado_a;