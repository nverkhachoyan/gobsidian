package transformers

import (
	"fmt"
	"gobsidian/internal/models"
	"regexp"
	"strings"
)

type CalloutTransformer struct {
	calloutRegex *regexp.Regexp
}

func NewCalloutTransformer(regex *regexp.Regexp) *CalloutTransformer {
	return &CalloutTransformer{
		calloutRegex: regex,
	}
}

func (ct *CalloutTransformer) Name() string {
	return "callout-transformer"
}

func (ct *CalloutTransformer) Transform(content string, note *models.ParsedNote, ctx *TransformContext) (string, error) {
	result := content

	return ct.calloutRegex.ReplaceAllStringFunc(result, func(match string) string {
		parts := ct.calloutRegex.FindStringSubmatch(match)
		if len(parts) < 3 {
			return match
		}

		calloutType := strings.ToLower(parts[1])
		rawContent := parts[2]

		lines := strings.Split(rawContent, "\n")
		var processedLines []string

		for _, line := range lines {
			if after, ok := strings.CutPrefix(line, ">"); ok {
				processedLine := after
				if after, ok := strings.CutPrefix(processedLine, " "); ok {
					processedLine = after
				}
				processedLines = append(processedLines, processedLine)
			} else if strings.TrimSpace(line) != "" {
				processedLines = append(processedLines, line)
			}
		}

		processedContent := strings.TrimSpace(strings.Join(processedLines, "\n"))

		return fmt.Sprintf(`<div class="obsidian-quote-%s"><div class="obsidian-quote-%s-title">%s</div>%s</div>`, calloutType, calloutType, calloutType, processedContent)
	}), nil
}
