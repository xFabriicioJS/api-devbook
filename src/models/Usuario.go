package models

import "time"

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
