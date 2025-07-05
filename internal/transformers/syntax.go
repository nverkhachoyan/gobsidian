package transformers

import (
	"fmt"
	"gobsidian/internal/models"
	"regexp"
	"strings"

	"github.com/alecthomas/chroma/v2/quick"
)

type SyntaxHighlighter struct {
	enabled        bool
	codeBlockRegex *regexp.Regexp
}

func NewSyntaxHighlighter(enabled bool) *SyntaxHighlighter {
	return &SyntaxHighlighter{
		enabled:        enabled,
		codeBlockRegex: regexp.MustCompile("```(\\w+)\\n([\\s\\S]*?)```"),
	}
}

func (s *SyntaxHighlighter) Name() string {
	return "syntax-highlighter"
}

func (s *SyntaxHighlighter) Transform(content string, note *models.ParsedNote, ctx *TransformContext) (string, error) {
	if !s.enabled {
		return content, nil
	}

	return s.codeBlockRegex.ReplaceAllStringFunc(content, func(match string) string {
		parts := s.codeBlockRegex.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}

		language := parts[1]
		code := parts[2]

		return s.highlightCode(code, language, ctx.IsDarkMode)

	}), nil
}

func (*SyntaxHighlighter) highlightCode(code, language string, isDarkMode bool) string {
	var buf strings.Builder

	theme := "github"
	if isDarkMode {
		theme = "monokai"
	}

	err := quick.Highlight(&buf, code, language, "html", theme)

	if err != nil {
		fmt.Printf("PROBLEMO: %v", err)
		return "<pre><code>" + code + "</code></pre>"
	}

	return buf.String()
}
