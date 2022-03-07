package models

import (
	"errors"
	"net/http"
	"strings"
	"time"
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
	user.format()

	return nil
}

// validate makes the user fields validation
func (user *User) validate(httpMethod string) error {
	if user.Name == "" {
		return errors.New("NAME is required")
	}

	if user.Login == "" {
		return errors.New("LOGIN is required")
	}

	if user.Email == "" {
		return errors.New("E-MAIL is required")
	}

	if user.Password == "" && httpMethod == http.MethodPost {
		return errors.New("PASSWORD is required")
	}

	return nil
}

// format makes the user fiels formatation
func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)

	user.Login = strings.TrimSpace(user.Name)
	user.Login = strings.ToLower(user.Login)

	user.Email = strings.TrimSpace(user.Name)
	user.Password = strings.TrimSpace(user.Name)
}
