package repositories

import (
	"api/src/models"
	"database/sql"
)

// user represent an user repository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository create and returns a user repository
func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

// Create insert an user in database
func (repository userRepository) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO users (name, login, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Login, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}
