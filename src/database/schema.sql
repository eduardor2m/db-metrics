CREATE TABLE IF NOT EXISTS "user" (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    preferences VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS book (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    genre VARCHAR(255) NOT NULL,
    synopsis TEXT
);

CREATE TABLE IF NOT EXISTS review (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES "user"(id),
    book_id UUID REFERENCES book(id),
    rating FLOAT NOT NULL,
    comment TEXT
);