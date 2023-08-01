package repositories

import (
	"api/src/models"
	"database/sql"
)

// Users represents a users repository
type Users struct {
	db *sql.DB
}

// NewUsersRepository creates a new Users repository
func NewUsersRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create inserts a new user in the database
func (u Users) Create(user models.User) (uint64, error) {
	return 0, nil
}
