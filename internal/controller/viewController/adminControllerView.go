package viewController

import (
	"Shiro/internal/service"
	"github.com/kataras/iris/v12"
)

type AdminController struct {
	userService *service.UserService
	postService *service.PostService
}

func NewAdminController(userService *service.UserService, postService *service.PostService) *AdminController {
	return &AdminController{
		userService: userService,
		postService: postService,
	}
}

func (c *AdminController) RegisterRoutes(app *iris.Application) {
	user := app.Party("/admin")
	{
		user.Get("/register", c.Register)
		user.Get("/login", c.Login)
		user.Get("/dashboard", c.Dashboard)
		user.Get("/postList", c.PostList)
		user.Get("/newPost", c.newPost)
	}
}

func (c *AdminController) newPost(ctx iris.Context) {
	ctx.ViewData("title", "New Post")
	err := ctx.View("/admin/newPost.jet")
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}

func (c *AdminController) PostList(ctx iris.Context) {
	posts, err := c.postService.GetAllPosts()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Error loading posts"})
		return
	}
	ctx.ViewData("posts", posts)
	ctx.ViewData("title", "Post List")
	err = ctx.View("/admin/postList.jet")
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}

func (c *AdminController) Dashboard(ctx iris.Context) {
	ctx.ViewData("title", "Dashboard")
	err := ctx.View("/admin/dashboard.jet")
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}

func (c *AdminController) Register(ctx iris.Context) {
	ctx.ViewData("title", "Register")
	err := ctx.View("/admin/register.jet")
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}

func (c *AdminController) Login(ctx iris.Context) {
	ctx.ViewData("title", "Login")
	err := ctx.View("/admin/login.jet")
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}
