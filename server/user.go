package main

import (
	"errors"

	"gorm.io/gorm"
)

const (
	MinPasswordLen = 8
)

type User struct {
	ID           int
	Name         string
	Email        string `gorm:"unique"`
	PasswordHash string
}

func hashPassword(password string) string {
	// TODO: implement hashing
	return password
}

func NewUser(name, email, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	if email == "" {
		return nil, errors.New("email is empty")
	}
	if len(password) <= MinPasswordLen {
		return nil, errors.New("password is empty or less than minimum required length")
	}
	u := &User{
		Name:         name,
		Email:        email,
		PasswordHash: hashPassword(password),
	}
	return u, nil
}
func (u *User) Save(db *gorm.DB) error {
	result := db.Create(u)
	return result.Error
}

func (u *User) SetPassword(password string) *User {
	hash := hashPassword(password)
	u.PasswordHash = hash
	return u
}
