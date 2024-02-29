package viewController

import (
	"Shiro/internal/repository"
	"Shiro/internal/service"
	"database/sql"
	"github.com/kataras/iris/v12"
)

func Register(app *iris.Application, db *sql.DB) {
	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postController := NewPostController(postService)
	postController.RegisterRoutes(app)

	pageController := NewIndexController(postService)
	pageController.RegisterRoutes(app)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	adminController := NewAdminController(userService, postService)
	adminController.RegisterRoutes(app)
}
