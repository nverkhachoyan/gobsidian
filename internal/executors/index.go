package executors

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type IndexExecutor struct {
	Config config.Config
}

func NewIndexExecutor(c config.Config) (*IndexExecutor, error) {
	return &IndexExecutor{Config: c}, nil
}

func (ie *IndexExecutor) ExecuteIndexPage(posts []*models.BlogPost, tags []models.Tag, fileTree *models.Folder) error {
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
	}{
		SiteTitle:    ie.Config.SiteTitle,
		SiteSubtitle: ie.Config.SiteSubtitle,
		Posts:        posts,
		Tags:         tags,
		CurrentYear:  time.Now().Year(),
		FileTree:     fileTree,
	}

	filePath := filepath.Join(ie.Config.OutputDirectory, "index.html")
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create index file %s: %w", filePath, err)
	}
	defer outFile.Close()

	if err := ie.Config.Templates.ExecuteTemplate(outFile, "index.html", data); err != nil {
		return fmt.Errorf("failed to execute index template: %w", err)
	}
	return nil
}
