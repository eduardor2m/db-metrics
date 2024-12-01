package models

import "github.com/google/uuid"

type Review struct {
	id      uuid.UUID
	userID  uuid.UUID
	bookID  uuid.UUID
	rating  float32
	comment string
}

func NewReview(id uuid.UUID, userID uuid.UUID, bookID uuid.UUID, rating float32, comment string) *Review {
	return &Review{
		id:      id,
		userID:  userID,
		bookID:  bookID,
		rating:  rating,
		comment: comment,
	}
}

func (r *Review) ID() uuid.UUID {
	return r.id
}

func (r *Review) UserID() uuid.UUID {
	return r.userID
}

func (r *Review) BookID() uuid.UUID {
	return r.bookID
}

func (r *Review) Rating() float32 {
	return r.rating
}

func (r *Review) Comment() string {
	return r.comment
}
