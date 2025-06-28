package executors

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
	"time"
)

type PostExecutor struct {
	Config config.Config
}

func NewPostExecutor(c config.Config) (*PostExecutor, error) {
	return &PostExecutor{Config: c}, nil
}

func (pe *PostExecutor) ExecutePostPage(p models.BlogPost) error {
	outputDir := pe.Config.OutputDirectory

	data := struct {
		models.BlogPost
		SiteTitle    string
		SiteSubtitle string
		BaseURL      string
		CurrentYear  int
	}{
		BlogPost:     p,
		SiteTitle:    pe.Config.SiteTitle,
		SiteSubtitle: pe.Config.SiteSubtitle,
		BaseURL:      pe.Config.BaseURL,
		CurrentYear:  time.Now().Year(),
	}

	filePath := filepath.Join(outputDir, p.FileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", filePath, err)
	}
	defer outFile.Close()

	if err := pe.Config.Templates.ExecuteTemplate(outFile, "post.html", data); err != nil {
		return fmt.Errorf("failed to execute post template: %w", err)
	}
	return nil
}
