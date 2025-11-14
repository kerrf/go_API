package models

import (
	"errors"

	"example.com/RestAPI/db"
	"example.com/RestAPI/utils"
)

type User struct {
	UserID   int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	results, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := results.LastInsertId()
	if err != nil {
		return err
	}

	u.UserID = id
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.UserID, &retrievedPassword)
	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsCorrect := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsCorrect {
		return errors.New("Credentials invalid")
	}

	return nil
}
