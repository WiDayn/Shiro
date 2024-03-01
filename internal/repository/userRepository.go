package repository

import (
	"database/sql"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo *UserRepository) CreateUser(username, password, email string) error {
	_, err := repo.Db.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", username, password, email)
	return err
}

func (repo *UserRepository) GetUserByUsername(username string) (string, error) {
	var password string
	err := repo.Db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&password)
	if err != nil {
		return "", err
	}
	return password, nil
}
