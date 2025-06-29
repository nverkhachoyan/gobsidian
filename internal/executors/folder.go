package executors

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
	"time"
)

func ExecuteFolderPage(
	cfg config.Config,
	folder *models.Folder,
	allTags []models.Tag,
) error {
	// Don't create a page for the root folder, it's handled by index.html
	if folder.Path == "" {
		return nil
	}

	outputDir := filepath.Join(cfg.OutputDirectory, folder.Path)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create folder directory %s: %w", outputDir, err)
	}

	data := struct {
		SiteTitle   string
		Folder      *models.Folder
		AllTags     []models.Tag
		CurrentYear int
	}{
		SiteTitle:   cfg.SiteTitle,
		Folder:      folder,
		AllTags:     allTags,
		CurrentYear: time.Now().Year(),
	}

	filePath := filepath.Join(outputDir, "index.html")
	outFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create folder index file %s: %w", filePath, err)
	}
	defer outFile.Close()

	if err := cfg.Templates.ExecuteTemplate(outFile, "folder.html", data); err != nil {
		return fmt.Errorf("failed to execute folder template for %s: %w", folder.Name, err)
	}
	return nil
}
