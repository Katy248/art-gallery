package users

import (
	"art-gallery-server/database"
	"art-gallery-server/models"
	"errors"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

const (
	MinPasswordLen = 8
)

type User struct {
	*models.BaseModel
	Name         string
	Email        string `gorm:"unique"`
	Description  string
	PasswordHash string
	IsAdmin      bool
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

// func (u *User) Save() error {
// 	result := db.Create(u)
// 	return result.Error
// }

func (u *User) SetPassword(password string) *User {
	hash := hashPassword(password)
	u.PasswordHash = hash
	return u
}
func (u *User) CheckPassword(password string) bool {
	return checkPassword(password, u.PasswordHash)
}

func GetUser(userId int) (User, error) {
	var user User
	result := database.Conn.Raw(`SELECT * FROM users WHERE id = ?`, userId).First(&user)
	return user, result.Error
}

func IsAdmin(userId int) (isAdmin bool) {
	result := database.Conn.Raw(rawIsAdminQuery, userId).First(&isAdmin)
	if result.Error != nil {
		log.Errorf("Failed to check if user is admin, returning default false value. Error: %s", result.Error)
		return false
	}
	return isAdmin
}

var rawIsAdminQuery = `
	SELECT is_admin FROM users WHERE id = ?
`
