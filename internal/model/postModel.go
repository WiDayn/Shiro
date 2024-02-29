package model

type Post struct {
	ID        int
	Title     string
	AuthorID  string
	Content   string
	CreatedAt string // 或使用time.Time
}
