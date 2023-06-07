package tool

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

const htmlTemplate = `<!doctype html>
<html>
<head>
<meta charset="UTF-8"><meta name="viewport" content="width=device-width initial-scale=1">
<title></title>
<link rel="stylesheet" href="https://unpkg.com/@highlightjs/cdn-assets@11.6.0/styles/atom-one-dark.min.css">
<script src="https://unpkg.com/@highlightjs/cdn-assets@11.6.0/highlight.min.js"></script>
<script src="https://unpkg.com/@highlightjs/cdn-assets@11.6.0/languages/go.min.js"></script>
</head>

<body>
%s
</body>
<script>
  hljs.highlightAll();
</script>
</html>`

func MarkdownToHtml(content string) string {
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		//goldmark.WithParserOptions(parser.WithAutoHeadingID()),
		//goldmark.WithRendererOptions(html.WithUnsafe()),
	)
	var buf bytes.Buffer
	err := md.Convert([]byte(content), &buf)
	if err != nil {
		return ""
	}

	return fmt.Sprintf(htmlTemplate, buf.String())
}
