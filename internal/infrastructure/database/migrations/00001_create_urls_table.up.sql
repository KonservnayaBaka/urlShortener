CREATE TABLE if NOT EXISTS urls (
    id INTEGER PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_url TEXT NOT NULL UNIQUE,
    created_at timestamp    not null default now()
);