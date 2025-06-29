package executors

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
	"time"
)

func ExecuteTagPage(cfg config.Config, tag models.Tag, posts []*models.BlogPost) error {
	// Create the tag directory for the tag page
	outputDir := filepath.Join(cfg.OutputDirectory, "tag", tag.Slug)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create tag directory %s: %w", outputDir, err)
	}

	data := struct {
		SiteTitle   string
		Tag         models.Tag
		Posts       []*models.BlogPost
		CurrentYear int
	}{
		SiteTitle:   cfg.SiteTitle,
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

	if err := cfg.Templates.ExecuteTemplate(outFile, "tag.html", data); err != nil {
		return fmt.Errorf("failed to execute tag template: %w", err)
	}
	return nil
}
