-- +goose Up
CREATE TABLE cientificos (
    dni varchar(10) PRIMARY KEY NOT NULL ,
    nomApels varchar(100) NOT NULL
);

-- +goose Down
DROP TABLE cientificos;