-- +migrate Up
CREATE TABLE container (
                       id integer PRIMARY KEY,
                       owner_address BYTEA,
                       tag TEXT,
                       recipient BYTEA[],
                       container TEXT,
);

-- +migrate Down
DROP TABLE container;
