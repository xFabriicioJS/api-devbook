package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// Find finds users by name or nick
func (repository Users) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	// Query for DB
	rows, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Find a single user, using id as parameter
func (repository Users) FindById(id uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"select id, name, nick, email, createdAt from users where id = ?",
		id,
	)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// Updating a single user, using id as parameter
func (repository Users) Update(userId uint64, user models.User) error {
	statement, err := repository.db.Prepare("update users set nome = ?, nick = ?, email = ?,where = ?")

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(
		user.Name,
		user.Nick,
		user.Email,
		userId,
	); err != nil {
		return err
	}

	return nil
}
