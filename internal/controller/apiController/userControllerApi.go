package apiController

import (
	"Shiro/internal/service"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) RegisterRoutes(app *iris.Application) {
	api := app.Party("/api/user")
	{
		api.Post("/register", c.Register)
		api.Post("/login", c.Login)
	}
}

func (c *UserController) Register(ctx iris.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	if err := ctx.ReadJSON(&request); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Invalid request"})
		return
	}

	err := c.userService.Register(request.Username, request.Password, request.Email)
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Failed to register user"})
		return
	}

	ctx.JSON(iris.Map{"message": "User registered successfully"})
}

func (c *UserController) Login(ctx iris.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ReadJSON(&request); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"message": "Invalid request"})
		return
	}

	success, err := c.userService.Login(request.Username, request.Password)
	if err != nil || !success {
		ctx.StatusCode(iris.StatusUnauthorized)
		ctx.JSON(iris.Map{"message": "Login failed"})
		return
	}

	ctx.JSON(iris.Map{"message": "Login successful"})
}
