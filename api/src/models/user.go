package models

import (
	"api/src/security"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represent an user in the app
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Login     string    `json:"login,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare validate and format the received user
func (user *User) Prepare(httpMethod string) error {
	if err := user.validate(httpMethod); err != nil {
		return err
	}

	user.format(httpMethod)

	return nil
}

// validate makes the user fields validation
func (user *User) validate(httpMethod string) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Login == "" {
		return errors.New("login is required")
	}

	if user.Email == "" {
		return errors.New("e-mail is required")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("e-mail with invalid format")
	}

	if user.Password == "" && httpMethod == http.MethodPost {
		return errors.New("PASSWORD is required")
	}

	return nil
}

// format makes the user fiels formatation
func (user *User) format(httpMethod string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Name = strings.Title(user.Name)

	user.Login = strings.TrimSpace(user.Name)
	user.Login = strings.ToLower(user.Login)

	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Name)

	if httpMethod == http.MethodPost {
		hashPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashPassword)
	}

	return nil
}
