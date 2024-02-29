package viewController

import (
	"Shiro/internal/service"
	"github.com/kataras/iris/v12"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{postService: postService}
}

func (c *PostController) RegisterRoutes(app *iris.Application) {
	app.Get("/posts", c.ShowPosts)
}

func (c *PostController) ShowPosts(ctx iris.Context) {
	posts, err := c.postService.GetAllPosts()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"message": "Error loading posts"})
		return
	}
	ctx.ViewData("posts", posts)
	ctx.ViewData("title", "Home Page")
	err = ctx.View("index.jet")
	if err != nil {
		ctx.StopWithText(iris.StatusInternalServerError, "Templates not rendered!")
	}
}
