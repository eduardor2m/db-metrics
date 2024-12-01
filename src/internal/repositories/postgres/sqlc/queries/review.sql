-- name: CreateReview :exec

INSERT INTO review (id, user_id, book_id, rating, comment) VALUES ($1, $2, $3, $4, $5);

-- name: ListReviews :many

SELECT * FROM review;