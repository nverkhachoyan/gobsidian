package transformers

import (
	"fmt"
	"regexp"

	"gobsidian/internal/models"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
)

type FootnoteEmbeddedTransformer struct {
	footnotesRegex *regexp.Regexp
	logger         *log.Logger
}

func NewFootnoteEmbeddedTransformer(
	footnotesRegex *regexp.Regexp,
	logger *log.Logger,
) *FootnoteEmbeddedTransformer {
	return &FootnoteEmbeddedTransformer{
		footnotesRegex: footnotesRegex,
		logger:         logger,
	}
}

func (ft *FootnoteEmbeddedTransformer) Name() string {
	return "footnote-embedded"
}

func (ft *FootnoteEmbeddedTransformer) Transform(
	body string,
	note *models.ParsedNote,
	ctx *TransformContext,
) (string, error) {
	if note.EmbeddedFootnotes != nil || len(note.EmbeddedFootnotes) > 0 {
		return body, nil
	}

	count := 1

	newBody := ft.footnotesRegex.ReplaceAllStringFunc(body, func(match string) string {
		parts := ft.footnotesRegex.FindStringSubmatch(match)
		footNoteNumber := parts[1]
		footnoteID := uuid.New().String()
		note.EmbeddedFootnotes = append(note.EmbeddedFootnotes, models.Footnote{
			ID:     footnoteID,
			Number: count,
			Text:   footNoteNumber,
		})
		result := fmt.Sprintf("<a href=\"#fn-%s\" class=\"footnote-ref\" id=\"fnref-%s\" role=\"doc-noteref\"><sup>[%d]</sup></a>", footnoteID, footnoteID, count)
		count++
		return result
	})

	return newBody, nil
}
