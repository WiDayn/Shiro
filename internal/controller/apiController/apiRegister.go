package apiController

import (
	"Shiro/internal/repository"
	"Shiro/internal/service"
	"database/sql"
	"github.com/kataras/iris/v12"
)

func Register(app *iris.Application, db *sql.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := NewUserController(userService)
	userController.RegisterRoutes(app)
}
