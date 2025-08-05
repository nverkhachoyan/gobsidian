package transformers

import (
	"fmt"
	"gobsidian/internal/models"
	"regexp"
	"strings"

	"github.com/alecthomas/chroma/v2"
	"github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/lexers"
	"github.com/alecthomas/chroma/v2/styles"
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

func (s *SyntaxHighlighter) Transform(content string, node *models.VaultNode, ctx *TransformContext) (string, error) {
	if node.GetNoteType() != models.NoteTypeMarkdown {
		return content, nil
	}

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

		if language == "mermaid" {
			return fmt.Sprintf("<div class='mermaid'>%s</div>", code)
		}

		return s.highlightCode(code, language)

	}), nil
}

func (*SyntaxHighlighter) highlightCode(code, language string) string {
	var buf strings.Builder
	lexer := lexers.Get(language)
	if lexer == nil {
		lexer = lexers.Analyse(code)
	}
	if lexer == nil {
		lexer = lexers.Fallback
	}
	lexer = chroma.Coalesce(lexer)

	formatter := html.New(
		html.WithClasses(true),
		html.WithLineNumbers(true),
		html.LineNumbersInTable(false),
	)

	iterator, err := lexer.Tokenise(nil, code)
	if err != nil {
		fmt.Printf("tokenise error: %v", err)
		return "<pre><code>" + code + "</code></pre>"
	}

	err = formatter.Format(&buf, styles.Fallback, iterator)
	if err != nil {
		fmt.Printf("format error: %v", err)
		return "<pre><code>" + code + "</code></pre>"
	}

	return buf.String()
}
