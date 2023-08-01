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
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}
