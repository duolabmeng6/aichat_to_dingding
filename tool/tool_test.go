package tool

import (
	"github.com/duolabmeng6/goefun/ecore"
	"testing"
)

func Test_MarkdownToHtml(t *testing.T) {
	txt := ecore.E读入文本("1.txt")
	html := MarkdownToHtml(txt)
	ecore.E写到文件("1.html", []byte(html))
}
