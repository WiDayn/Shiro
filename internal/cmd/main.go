package main

import (
	"Shiro/internal/controller/apiController"
	"Shiro/internal/controller/viewController"
	"Shiro/internal/data"
	"Shiro/internal/middleware"
	"Shiro/internal/template"
	"Shiro/internal/theme"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris/v12"
)

func main() {
	// 启动Iris
	app := iris.New()

	// 初始化数据库
	db := data.SetupDB(app)
	defer db.Close()

	// 检查主题完整性
	theme.IntegrityCheck("./themes/BlackOrWhite/template")
	template.LoadDir(app, "./themes/BlackOrWhite/template")

	app.Use(middleware.JWTMiddleware)

	apiController.Register(app, db)
	viewController.Register(app, db)

	app.HandleDir("/static", "./views/BlackOrWhite/static")

	app.Listen(":8080")
}
