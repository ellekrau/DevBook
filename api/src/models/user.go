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
	Nickname  string    `json:"nickname,omitempty"`
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

	if user.Nickname == "" {
		return errors.New("nickname is required")
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

	user.Nickname = strings.TrimSpace(user.Name)
	user.Nickname = strings.ToLower(user.Nickname)

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
