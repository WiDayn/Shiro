package viewController

import (
	"Shiro/internal/service"
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
	ctx.ViewData("posts", posts)
	ctx.ViewData("title", "Home Page")
	ctx.View("index.jet")
}
