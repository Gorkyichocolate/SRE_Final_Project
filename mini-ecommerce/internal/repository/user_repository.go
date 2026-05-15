package repository

import (
	"sre/internal/database"
	"sre/internal/model"
)

func CreateUser(
	username,
	email,
	passwordHash string,
) error {

	_, err := database.DB.Exec(
		`
		INSERT INTO users(username, email, password_hash)
		VALUES($1, $2, $3)
		`,
		username,
		email,
		passwordHash,
	)

	return err
}

func GetUserByEmail(email string) (*model.User, error) {

	var user model.User

	err := database.DB.QueryRow(
		`
		SELECT id, username, email, password_hash
		FROM users
		WHERE email=$1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
