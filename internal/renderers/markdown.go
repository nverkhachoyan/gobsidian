package renderers

import (
	"github.com/gomarkdown/markdown/html"
)

var rendererConfig = html.RendererOptions{
	Flags: html.CommonFlags | html.HrefTargetBlank,
}

func NewMarkdownRenderer() *html.Renderer {
	return html.NewRenderer(rendererConfig)
}
