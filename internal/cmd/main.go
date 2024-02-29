package main

import (
	"Shiro/internal/controller/apiController"
	"Shiro/internal/controller/viewController"
	"Shiro/internal/data"
	"Shiro/internal/template"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
)

func main() {
	// 启动Iris
	app := iris.New()

	// 初始化数据库
	db := data.SetupDB(app)
	defer db.Close()
	template.LoadDir(app, "./views/BlackOrWhite")

	apiController.Register(app, db)
	viewController.Register(app, db)

	app.HandleDir("/static", "./views/BlackOrWhite/static")

	app.Listen(":8080")
}
