-- +migrate Up
CREATE TABLE containers (
                       id  SERIAL PRIMARY KEY not NULL ,
                       owner BYTEA not NULL,
                       tag TEXT not NULL,
                       recipient TEXT[] not NULL,
                       container BYTEA not NULL
);

-- +migrate Down
DROP TABLE containers;
