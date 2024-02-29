package repository

import (
	"Shiro/internal/model"
	"database/sql"
)

type PostRepository struct {
	Db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{Db: db}
}

func (repo *PostRepository) GetAllPosts() ([]model.Post, error) {
	var posts []model.Post
	rows, err := repo.Db.Query("SELECT id, author_id, title, content, created_at FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.AuthorID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
