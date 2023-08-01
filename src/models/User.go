package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a user in the system
type User struct {
	// "omitempty" means that if the field is empty, it will not be displayed in the JSON

	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare will call the methods to validate and format the user struct  received
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()

	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("The field 'name' is required and cannot be empty")
	}
	if user.Nick == "" {
		return errors.New("The field 'nick' is required and cannot be empty")
	}
	if user.Email == "" {
		return errors.New("The field 'email' is required and cannot be empty")
	}
	if user.Password == "" {
		return errors.New("The field 'password' is required and cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
