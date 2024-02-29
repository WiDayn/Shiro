package routes

import (
	"database/sql"
	"github.com/kataras/iris/v12"
)

func SetupRoutes(app *iris.Application, db *sql.DB) {
	indexRoutes(app, db)
}
