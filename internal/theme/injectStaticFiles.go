package theme

import (
	"Shiro/internal/config"
	"Shiro/internal/tools"
	"github.com/kataras/iris/v12"
	"os"
	"strings"
)

func resolveTemplateInheritanceChain(templatePath string) []string {
	themePath := "themes/" + config.Application.ThemeConfig.ThemeDir + "/template"

	parts := strings.Split(templatePath, "/")
	var chain []string
	nowPath := ""
	for i, part := range parts {
		nowPath = nowPath + part
		if i != len(parts)-1 {
			nowPath += "/"
		}
		if fileExists(themePath + nowPath + "base.jet") {
			chain = append(chain, nowPath+"base.jet")
		}
	}
	if parts[len(parts)-1] != "base.jet" {
		if fileExists(themePath + nowPath) {
			chain = append(chain, nowPath)
		}
	}
	return chain
}

// fileExists checks if a file exists at the given path.
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// embedStaticFile 读取指定路径的文件内容并返回
func embedStaticFile(path string) (string, error) {
	// 读取文件内容
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// InjectStaticFiles 自动根据Jet模板路径注入对应的CSS和JS文件
func InjectStaticFiles(ctx iris.Context, templatePath string) (string, string, error) {
	chain := resolveTemplateInheritanceChain(templatePath)
	cssRootPath := "themes/" + config.Application.ThemeConfig.ThemeDir + "/css"
	jsRootPath := "themes/" + config.Application.ThemeConfig.ThemeDir + "/js"

	var combinedCSS, combinedJS string
	for _, tmplPath := range chain {
		cssPath := strings.Replace(cssRootPath+tmplPath, ".jet", ".css", 1)
		jsPath := strings.Replace(jsRootPath+tmplPath, ".jet", ".js", 1)

		if cssContent, err := embedStaticFile(cssPath); err == nil {
			combinedCSS += cssContent
		}

		if jsContent, err := embedStaticFile(jsPath); err == nil {
			combinedJS += jsContent
		}
	}
	compressedCSS, _ := tools.MinifyContent("text/css", combinedCSS)
	compressedJS, _ := tools.MinifyContent("application/javascript", combinedJS)
	ctx.ViewData("cssContent", "<style>"+compressedCSS+"</style>")
	ctx.ViewData("jsContent", "<script>"+compressedJS+"</script>")
	return combinedCSS, combinedJS, nil
}
