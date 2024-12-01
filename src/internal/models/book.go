package models

import "github.com/google/uuid"

type Book struct {
	id       uuid.UUID
	title    string
	author   string
	genre    string
	synopsis string
}

func NewBook(id uuid.UUID, title string, author string, genre string, synopsis string) *Book {
	return &Book{
		id:       id,
		title:    title,
		author:   author,
		genre:    genre,
		synopsis: synopsis,
	}
}

func (b *Book) ID() uuid.UUID {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Author() string {
	return b.author
}

func (b *Book) Genre() string {
	return b.genre
}

func (b *Book) Synopsis() string {
	return b.synopsis
}
