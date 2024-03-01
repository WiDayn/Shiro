package viewController

import (
	"Shiro/internal/middleware"
	"Shiro/internal/service"
	"Shiro/internal/theme"
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
	userNoAuth := app.Party("/admin")
	{
		userNoAuth.Get("/register", c.Register)
		userNoAuth.Get("/login", c.Login)
	}
	user := app.Party("/admin", middleware.JWTViewRequiredCheck)
	{
		user.Get("/dashboard", c.Dashboard)
		user.Get("/post-list", c.PostList)
		user.Get("/new-post", c.newPost)
	}
}

func (c *AdminController) newPost(ctx iris.Context) {
	templatePath := "/admin/panel/new-post.jet" // 模板文件名
	// 自动注入CSS和JS
	theme.InjectStaticFiles(ctx, templatePath)
	ctx.ViewData("title", "New Post")
	err := ctx.View(templatePath)
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
	templatePath := "/admin/panel/post-list.jet" // 模板文件名
	// 自动注入CSS和JS
	theme.InjectStaticFiles(ctx, templatePath)
	ctx.ViewData("posts", posts)
	ctx.ViewData("title", "Post List")
	err = ctx.View(templatePath)
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}

func (c *AdminController) Dashboard(ctx iris.Context) {
	templatePath := "/admin/panel/dashboard.jet" // 模板文件名
	// 自动注入CSS和JS
	theme.InjectStaticFiles(ctx, templatePath)
	ctx.ViewData("title", "Dashboard")
	err := ctx.View(templatePath)
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}

func (c *AdminController) Register(ctx iris.Context) {
	templatePath := "/admin/login/register.jet" // 模板文件名
	// 自动注入CSS和JS
	theme.InjectStaticFiles(ctx, templatePath)
	ctx.ViewData("title", "Register")
	err := ctx.View(templatePath)
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}

func (c *AdminController) Login(ctx iris.Context) {
	templatePath := "/admin/login/login.jet" // 模板文件名
	// 自动注入CSS和JS
	theme.InjectStaticFiles(ctx, templatePath)
	ctx.ViewData("title", "Login")
	err := ctx.View(templatePath)
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}
