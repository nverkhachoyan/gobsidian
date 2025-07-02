package executors

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"gobsidian/internal/utils"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

func ExecutePostPage(cfg config.Config, p models.BlogPost, tags []models.Tag, fileTree *models.Folder, graph []byte) error {
	outputDir := cfg.OutputDirectory

	data := struct {
		models.BlogPost
		SiteTitle    string
		SiteSubtitle string
		BaseURL      string
		CurrentYear  int
		Tags         []models.Tag
		FileTree     *models.Folder
		Graph        template.JS
	}{
		BlogPost:     p,
		SiteTitle:    cfg.SiteTitle,
		SiteSubtitle: cfg.SiteSubtitle,
		BaseURL:      cfg.BaseURL,
		CurrentYear:  time.Now().Year(),
		Tags:         tags,
		FileTree:     fileTree,
		Graph:        template.JS(graph),
	}

	filePath := filepath.Join(outputDir, utils.Slugify(p.RelativePathWithoutName), p.FileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", filePath, err)
	}
	defer outFile.Close()

	if err := cfg.Templates.ExecuteTemplate(outFile, "post.html", data); err != nil {
		return fmt.Errorf("failed to execute post template: %w", err)
	}
	return nil
}
