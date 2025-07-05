package transformers

import (
	"fmt"
	"gobsidian/internal/models"
	"regexp"
)

type HashtagEnricher struct {
	hashtagRegex *regexp.Regexp
}

func NewHashtagEnricher(regex *regexp.Regexp) *HashtagEnricher {
	return &HashtagEnricher{
		hashtagRegex: regex,
	}
}

func (ip *HashtagEnricher) Name() string {
	return "hashtag-enricher"
}

func (he *HashtagEnricher) Transform(content string, note *models.ParsedNote, ctx *TransformContext) (string, error) {
	result := content
	return he.hashtagRegex.ReplaceAllStringFunc(result, func(match string) string {
		parts := he.hashtagRegex.FindStringSubmatch(match)
		tagName := parts[1]
		return fmt.Sprintf(`<a class="tag" href="/tag/%s">#%s</a>`, tagName, tagName)
	}), nil
}
