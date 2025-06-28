package executors

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
	"time"
)

type TagExecutor struct {
	Config config.Config
}

func NewTagExecutor(c config.Config) (*TagExecutor, error) {
	return &TagExecutor{Config: c}, nil
}

func (ce *TagExecutor) ExecuteTagPage(tag models.Tag, posts []*models.BlogPost) error {
	// Create the tag directory for the tag page
	outputDir := filepath.Join(ce.Config.OutputDirectory, "tag", tag.Slug)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create tag directory %s: %w", outputDir, err)
	}

	data := struct {
		SiteTitle   string
		Tag         models.Tag
		Posts       []*models.BlogPost
		CurrentYear int
	}{
		SiteTitle:   ce.Config.SiteTitle,
		Tag:         tag,
		Posts:       posts,
		CurrentYear: time.Now().Year(),
	}

	filePath := filepath.Join(outputDir, "index.html")
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create tag index file %s: %w", filePath, err)
	}
	defer outFile.Close()

	if err := ce.Config.Templates.ExecuteTemplate(outFile, "tag.html", data); err != nil {
		return fmt.Errorf("failed to execute tag template: %w", err)
	}
	return nil
}
