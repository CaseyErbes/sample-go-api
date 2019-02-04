-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE address (
    id              uuid PRIMARY KEY NOT NULL,
    firstname       varchar(128),
    lastname        varchar(128),
    email           varchar(256) NOT NULL,
    phonenumber     varchar(15) NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE address;
