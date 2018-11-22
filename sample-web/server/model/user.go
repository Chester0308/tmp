package model

import (
)

type User struct {
	UserID int
	Name string
}

// NewUser create new User model
func NewUser(UserID int, username string) (*User, error) {
	u := new(User)
	u.UserID = UserID
	u.Name = username
	return u, nil
}