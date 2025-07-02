package executors

import (
	"fmt"
	"gobsidian/internal/config"
	"html/template"
	"os"
	"path/filepath"
	"time"
)

func ExecuteNotFoundPage(cfg config.Config, graph []byte) error {
	outputDir := cfg.OutputDirectory

	data := struct {
		SiteTitle    string
		SiteSubtitle string
		BaseURL      string
		CurrentYear  int
		Graph        template.JS
	}{
		SiteTitle:    cfg.SiteTitle,
		SiteSubtitle: cfg.SiteSubtitle,
		BaseURL:      cfg.BaseURL,
		CurrentYear:  time.Now().Year(),
		Graph:        template.JS(graph),
	}

	outFile, err := os.Create(filepath.Join(outputDir, "404.html"))
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", "404.html", err)
	}
	defer outFile.Close()

	if err := cfg.Templates.ExecuteTemplate(outFile, "404.html", data); err != nil {
		return fmt.Errorf("failed to execute 404 template: %w", err)
	}
	return nil
}
