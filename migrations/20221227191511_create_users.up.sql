CREATE TABLE users(
    PRIMARY KEY(id),

    id BIGSERIAL,
    email VARCHAR NOT NULL UNIQUE,
    encrypted_password VARCHAR NOT NULL
);