package executors

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
)

func ExecutePreviewPage(cfg config.Config, p models.BlogPost) error {
	outputDir := cfg.OutputDirectory
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create previews directory: %w", err)
	}

	filePath := filepath.Join(outputDir, "previews", p.RelativePath, p.FileName)
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create output file %s: %w", filePath, err)
	}
	defer outFile.Close()

	if err := cfg.Templates.ExecuteTemplate(outFile, "preview.html", p); err != nil {
		return fmt.Errorf("failed to execute preview template: %w", err)
	}
	return nil
}
