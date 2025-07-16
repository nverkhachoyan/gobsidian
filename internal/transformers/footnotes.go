package transformers

import (
	"fmt"
	"regexp"

	"gobsidian/internal/models"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

type FootnoteTransformer struct {
	footnotesRegex *regexp.Regexp
	logger         *log.Logger
}

func NewFootnoteTransformer(
	footnotesRegex *regexp.Regexp,
	logger *log.Logger,
) *FootnoteTransformer {
	return &FootnoteTransformer{
		footnotesRegex: footnotesRegex,
		logger:         logger,
	}
}

func (ft *FootnoteTransformer) Name() string {
	return "footnote"
}

func (ft *FootnoteTransformer) Transform(
	body string,
	note *models.ParsedNote,
	ctx *TransformContext,
) (string, error) {
	if note.Footnotes != nil || len(note.Footnotes) > 0 {
		return body, nil
	}

	count := 1
	footnotes := []models.Footnote{}

	newBody := ft.footnotesRegex.ReplaceAllStringFunc(body, func(match string) string {
		parts := ft.footnotesRegex.FindStringSubmatch(match)
		footNoteNumber := parts[1]
		footnoteID := uuid.New().String()
		footnotes = append(footnotes, models.Footnote{
			ID:     footnoteID,
			Number: count,
			Text:   footNoteNumber,
		})
		result := fmt.Sprintf("<a href=\"#fn-%s\" class=\"footnote-ref\" id=\"fnref-%s\" role=\"doc-noteref\"><sup>[%d]</sup></a>", footnoteID, footnoteID, count)
		count++
		return result
	})

	note.Footnotes = footnotes

	return newBody, nil
}
