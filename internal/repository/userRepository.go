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

func (repo *UserRepository) GetUserByUsername(username string) (int, string, error) {
	var id int
	var password string
	err := repo.Db.QueryRow("SELECT id, password FROM users WHERE username = ?", username).Scan(&id, &password)
	if err != nil {
		return 0, "", err
	}
	return id, password, nil
}
