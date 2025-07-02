package executors

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"time"
)

func ExecuteIndexPage(
	cfg config.Config,
	posts []*models.BlogPost,
	tags []models.Tag,
	fileTree *models.Folder,
	graph []byte,
) error {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	data := struct {
		SiteTitle    string
		SiteSubtitle string
		Posts        []*models.BlogPost
		Tags         []models.Tag
		CurrentYear  int
		FileTree     *models.Folder
		Graph        template.JS
	}{
		SiteTitle:    cfg.SiteTitle,
		SiteSubtitle: cfg.SiteSubtitle,
		Posts:        posts,
		Tags:         tags,
		CurrentYear:  time.Now().Year(),
		FileTree:     fileTree,
		Graph:        template.JS(graph),
	}

	filePath := filepath.Join(cfg.OutputDirectory, "index.html")
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create index file %s: %w", filePath, err)
	}
	defer outFile.Close()

	if err := cfg.Templates.ExecuteTemplate(outFile, "index.html", data); err != nil {
		return fmt.Errorf("failed to execute index template: %w", err)
	}
	return nil
}
