package routes

import (
	"Shiro/internal/model"
	"database/sql"
	"github.com/kataras/iris/v12"
)

func indexRoutes(app *iris.Application, db *sql.DB) {
	app.Get("/", func(ctx iris.Context) {
		var posts []model.Post
		rows, err := db.Query("SELECT id, title, content, created_at FROM posts")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef("Error: %v", err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			var post model.Post
			if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
				ctx.StatusCode(iris.StatusInternalServerError)
				ctx.Writef("Error: %v", err)
				return
			}
			posts = append(posts, post)
		}
		ctx.View("index.jet", posts)
	})
}
