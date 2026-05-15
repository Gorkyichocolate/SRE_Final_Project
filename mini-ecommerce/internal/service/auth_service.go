package service

import (
	"errors"

	"sre/internal/repository"
	"sre/internal/utils"
)

func RegisterUser(
	username,
	email,
	password string,
) error {

	hash, err := utils.HashPassword(password)

	if err != nil {
		return err
	}

	return repository.CreateUser(
		username,
		email,
		hash,
	)
}

func LoginUser(
	email,
	password string,
) (string, error) {

	user, err := repository.GetUserByEmail(email)

	if err != nil {
		return "", err
	}

	valid := utils.CheckPassword(
		password,
		user.PasswordHash,
	)

	if !valid {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
