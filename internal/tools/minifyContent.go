package tools

import (
	"bytes"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/js"
)

// MinifyContent 使用tdewolff/minify库来压缩CSS和JS内容
func MinifyContent(contentType, content string) (string, error) {
	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("application/javascript", js.Minify)

	var b bytes.Buffer
	if err := m.Minify(contentType, &b, bytes.NewReader([]byte(content))); err != nil {
		return "", err
	}
	return b.String(), nil
}
