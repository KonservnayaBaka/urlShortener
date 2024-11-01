CREATE TABLE if NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    original_url VARCHAR(255) NOT NULL,
    short_url VARCHAR(255) NOT NULL UNIQUE,
    short_url_with_domain VARCHAR(255) NOT NULL UNIQUE,
    created_at timestamp not null default now()
);

CREATE TABLE if NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) NOT NULL,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL UNIQUE,
)