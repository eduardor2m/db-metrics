package models

import "github.com/google/uuid"

type User struct {
	id          uuid.UUID
	name        string
	email       string
	preferences string
}

func NewUser(id uuid.UUID, name string, email string, preferences string) *User {
	return &User{
		id:          id,
		name:        name,
		email:       email,
		preferences: preferences,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Preferences() string {
	return u.preferences
}
