package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

// GetUsers get users in database filtering by name or login
func (repository userRepository) GetUsers(nameOrLogin string) (users []models.User, err error) {
	nameOrLogin = fmt.Sprintf("%%%s%%", nameOrLogin) //%nameOrLogin%

	rows, err := repository.db.Query(
		"SELECT id, name, login, email, createdAt FROM users WHERE name LIKE ? OR login LIKE ?", nameOrLogin, nameOrLogin)
	if err != nil {
		return
	}

	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Login, &user.Email, &user.CreatedAt); err != nil {
			return
		}

		users = append(users, user)
	}
	defer rows.Close()

	return
}

// GetUserById use one ID to get an user from database
func (repository userRepository) GetUserById(userID uint64) (user models.User, err error) {
	row, err := repository.db.Query(
		"SELECT id, name, login, email, createdAt FROM users WHERE id = ?", userID)
	if err != nil {
		return
	}
	defer row.Close()

	user = models.User{}
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Login, &user.Email, &user.CreatedAt); err != nil {
			return
		}
	}

	return
}

// UpdateUser update an user in database
func (repository userRepository) UpdateUser(userID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"UPDATE users SET name = ?, login = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Login, user.Email, userID); err != nil {
		return err
	}

	return err
}
