package models

import (
	"time"
	"database/sql"
	"log"
)

type User struct {
	UserId int `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	IsSessionAlive bool `json:"is_session_alive"`
	PasswordDigest string `json:"password_digest"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func NewUser() User {
	return User{}
}

func (s User) Insert(db *sql.DB, user User) error {
	stmt, err := db.Prepare(`INSERT INTO users SET name=?, email=?, password_digest=?, is_session_alive=?`)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	_, err = stmt.Exec(user.Name, user.Email, user.PasswordDigest, true)
	if err != nil {
		log.Fatalf("error: %s", err)
		return err
	}
	return nil
}

func (s User) GetByUserId(db *sql.DB, id int64) (User, error) {
	var user User
	err := db.QueryRow(`SELECT * FROM users WHERE user_id=? limit 1`, id).Scan(
		&user.UserId,
		&user.Name,
		&user.Email,
		&user.IsSessionAlive,
		&user.PasswordDigest,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
