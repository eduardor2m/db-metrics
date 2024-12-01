CREATE TABLE IF NOT EXISTS review (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES "user"(id),
    book_id UUID REFERENCES book(id),
    rating FLOAT NOT NULL,
    comment TEXT
);