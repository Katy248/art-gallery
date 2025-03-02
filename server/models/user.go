package models

import (
	"errors"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
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
	PostSaves    []PostSave
}

func hashPassword(password string) string {
	cost := viper.GetInt("password.cost")
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		log.Fatalf("Failed hash password: %s", err)
		return ""
	}
	return string(hashed)
}
func checkPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
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
func (u *User) CheckPassword(password string) bool {
	// hashed := hashPassword(password)
	// log.Debugf("Compare passwords '%s' and '%s'", u.PasswordHash, hashed)
	// return u.PasswordHash == hashed
	return checkPassword(password, u.PasswordHash)
}
