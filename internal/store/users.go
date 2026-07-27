package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
)

type User struct {
	ID        int64  `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	val, er := json.Marshal(&user)
	if er != nil {
		log.Fatalf("Failed to encode JSON: %v", er)
	}

	// 2. Convert []byte to string
	jsonStr := string(val)
	log.Fatal(jsonStr)

	query := `
		INSERT INTO users (first_name, last_name, username, password, email)
		VALUES ($1, $2, $3, $4, $5) RETURNING id, first_name, last_name, email, created_at`

	err := s.db.QueryRowContext(
		ctx, query, user.Firstname, user.Lastname, user.Username, user.Password, user.Email,
	).Scan(
		&user.ID,
		&user.Firstname,
		&user.Lastname,
		&user.Email,
		&user.CreatedAt,
	)

	if err != nil {
		log.Fatal(err.Error())
		return err
	}

	return nil
}
