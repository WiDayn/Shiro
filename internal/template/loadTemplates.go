package template

import (
	"github.com/kataras/iris/v12"
)

// LoadDir 遍历并加载目录及子目录中的模板
func LoadDir(app *iris.Application, dir string) {
	app.RegisterView(iris.Jet(dir, ".jet").Reload(true))

	app.Logger().Info("The views have been successfully registered")
}
