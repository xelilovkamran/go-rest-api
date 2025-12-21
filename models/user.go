package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := db.DB.Exec(`INSERT INTO users (email, password) VALUES (?, ?)`,
		u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = int(id)
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	row := db.DB.QueryRow(`SELECT id, email, password FROM users WHERE email = ?`, email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(id int) (*User, error) {
	row := db.DB.QueryRow(`SELECT id, email, password FROM users WHERE id = ?`, id)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}