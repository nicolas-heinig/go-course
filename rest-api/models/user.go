package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password)
	VALUES ($1, $2)
	RETURNING id
	`

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	err = db.DB.QueryRow(query, u.Email, hashedPassword).Scan(&u.ID)

	if err != nil {
		return err
	}

	return err
}

func (u *User) ValidateCreds() error {
	query := "SELECT id, password FROM users WHERE email = $1"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials invalid!")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials invalid!")
	}

	return nil
}
