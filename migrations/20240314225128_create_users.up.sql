CREATE TABLE users (
    id bigserial not null primary key,
    email varchar not null unique,
    enctypted_password varchar not null
);