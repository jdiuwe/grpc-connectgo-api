CREATE TABLE IF NOT EXISTS users
(
    id              BIGSERIAL PRIMARY KEY,
    uuid            VARCHAR UNIQUE NOT NULL,
    firstname       VARCHAR        NOT NULL,
    lastname        VARCHAR        NOT NULL,
    email           VARCHAR UNIQUE NOT NULL,
    hashed_password VARCHAR        NOT NULL,
    verified        BOOLEAN                 DEFAULT FALSE,
    created         timestamptz    NOT NULL DEFAULT now()
);
