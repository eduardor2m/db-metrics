-- name: CreateBook :exec

INSERT INTO book (id, title, author, genre, synopsis) VALUES ($1, $2, $3, $4, $5);

-- name: ListBooks :many

SELECT * FROM book;