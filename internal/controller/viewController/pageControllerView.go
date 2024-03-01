package viewController

import (
	"Shiro/internal/service"
	"Shiro/internal/theme"
	"github.com/kataras/iris/v12"
)

type PageController struct {
	postService *service.PostService
}

func NewIndexController(postService *service.PostService) *PageController {
	return &PageController{postService: postService}
}

func (c *PageController) RegisterRoutes(app *iris.Application) {
	app.Get("/", c.ShowIndex)
}

func (c *PageController) ShowIndex(ctx iris.Context) {
	posts, err := c.postService.GetAllPosts()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Error loading posts"})
		return
	}
	templatePath := "/page/index.jet" // 模板文件名
	// 自动注入CSS和JS
	theme.InjectStaticFiles(ctx, templatePath)

	ctx.ViewData("username", ctx.Values().Get("username"))
	ctx.ViewData("posts", posts)
	ctx.ViewData("title", "Home Page")
	ctx.View(templatePath)
}
