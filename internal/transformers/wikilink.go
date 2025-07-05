package transformers

import (
	"bytes"
	"fmt"
	"gobsidian/internal/models"
	"gobsidian/internal/utils"
	"html/template"
	"regexp"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/gomarkdown/markdown"

	"github.com/gomarkdown/markdown/html"
	"github.com/google/uuid"
)

type WikilinkTransformer struct {
	logger           *log.Logger
	wikilinkRegex    *regexp.Regexp
	imageLinkRegex   *regexp.Regexp
	hashtagRegex     *regexp.Regexp
	markdownRenderer *html.Renderer
}

func NewWikilinkTransformer(
	wikilinkRegexp *regexp.Regexp,
	imageLinkRegex *regexp.Regexp,
	hashtagRegex *regexp.Regexp,
	logger *log.Logger,
	markdownRenderer *html.Renderer,
) *WikilinkTransformer {
	return &WikilinkTransformer{
		wikilinkRegex:    wikilinkRegexp,
		logger:           logger,
		imageLinkRegex:   imageLinkRegex,
		hashtagRegex:     hashtagRegex,
		markdownRenderer: markdownRenderer,
	}
}

func (wt *WikilinkTransformer) Name() string {
	return "wikilink-transformer"
}

func (wt *WikilinkTransformer) Transform(
	body string,
	note *models.ParsedNote,
	ctx *TransformContext,
) (string, error) {

	return wt.wikilinkRegex.ReplaceAllStringFunc(body, func(match string) string {
		parts := wt.wikilinkRegex.FindStringSubmatch(match)
		matchedLinkText := parts[1]
		pseudoName := matchedLinkText

		if len(parts) > 2 && parts[2] != "" {
			pseudoName = parts[2]
		}

		isEmbedded := strings.HasPrefix(match, "!")

		if targetNote, found := ctx.notesRepository.ResolveWikilink(matchedLinkText); found {
			ctx.notesRepository.AddBacklink(targetNote, note, pseudoName)

			if isEmbedded {
				return wt.handleEmbeddedPost(ctx.EmbeddedPosts, targetNote, ctx)
			}
			return fmt.Sprintf(`<a href="%s">%s</a>`, targetNote.URL, pseudoName)
		}

		return fmt.Sprintf(`<a href="/%s">%s</a>`, utils.Slugify(matchedLinkText)+".html", pseudoName)
	}), nil
}

func (wt *WikilinkTransformer) handleEmbeddedPost(
	embeddedPosts map[string]models.EmbeddedPost,
	embeddedPost *models.ParsedNote,
	ctx *TransformContext,
) string {
	var uniqueID string

	htmlContent := wt.renderPostContent(embeddedPost, ctx)
	uniqueID = uuid.New().String()
	embeddedPosts[uniqueID] = models.EmbeddedPost{
		ID:      uniqueID,
		Content: htmlContent,
	}
	return fmt.Sprintf(`<div class="embedded-post">
		<div class="embedded-post-header">
		<a href="%s"><h2>%s</h2></a>
		</div>
		%s</div>`, embeddedPost.URL, embeddedPost.Title, uniqueID)
}

func (wt *WikilinkTransformer) renderPostContent(
	post *models.ParsedNote,
	ctx *TransformContext,
) template.HTML {
	if renderedPost, ok := ctx.getCachedPost(post.URL); ok {
		return renderedPost
	}

	if ctx.isRenderingInProgress(post.URL) {
		wt.logger.Warn("Circular embed detected, returning placeholder/error for", "title", post.Title)
		return template.HTML(fmt.Sprintf(`<div class="embedded-post-circular-error">Circular Embed: %s</div>`, post.Title))
	}

	ctx.setRenderingInProgress(post.URL, true)

	pipeline := NewPipeline(
		NewImageProcessor(wt.imageLinkRegex),
		NewWikilinkTransformer(wt.wikilinkRegex, wt.imageLinkRegex, wt.hashtagRegex, wt.logger, wt.markdownRenderer),
		NewHashtagEnricher(wt.hashtagRegex),
	)

	result, err := pipeline.Transform(string(post.RawBody), post, ctx)
	if err != nil {
		wt.logger.Error("pipeline transform failed", "error", err)
		return template.HTML(fmt.Sprintf(`<div class="error">Failed to render: %s</div>`, post.Title))
	}

	finalHTMLBytes := markdown.ToHTML([]byte(result), nil, wt.markdownRenderer)

	for _, embeddedPost := range ctx.EmbeddedPosts {
		finalHTMLBytes = bytes.Replace(finalHTMLBytes, []byte(embeddedPost.ID), []byte(embeddedPost.Content), 1)
	}

	ctx.setCachedPost(post.URL, template.HTML(finalHTMLBytes))
	ctx.setRenderingInProgress(post.URL, false)

	return template.HTML(finalHTMLBytes)
}
