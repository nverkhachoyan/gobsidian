package builders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gobsidian/internal/crawler"
	"gobsidian/internal/models"

	"github.com/charmbracelet/log"
)

type SearchIndexBuilder struct {
	Logger             *log.Logger
	outputDirectory    string
	generatedDirectory string
	fileName           string
	searchIndex        *models.SearchIndex
}

func NewSearchIndexBuilder(logger *log.Logger, outputDirectory string, generatedDirectory string, fileName string) *SearchIndexBuilder {
	return &SearchIndexBuilder{
		Logger:             logger,
		outputDirectory:    outputDirectory,
		generatedDirectory: generatedDirectory,
		fileName:           fileName,
		searchIndex: &models.SearchIndex{
			Notes: make([]*models.SearchIndexNote, 0),
		},
	}
}

func (s *SearchIndexBuilder) Build(fileIndex map[string]*crawler.VaultNode) time.Duration {
	start := time.Now()
	s.searchIndex = &models.SearchIndex{
		Notes: s.searchIndex.Notes,
	}

	for _, node := range fileIndex {

		shortBody := node.HTML
		if len(shortBody) > 100 {
			shortBody = shortBody[:100] + "..."
		}

		s.searchIndex.Notes = append(s.searchIndex.Notes, &models.SearchIndexNote{
			ID:    node.ID,
			URL:   node.URL,
			Title: node.Title,
			Body:  string(shortBody),
			Tags:  node.Tags,
		})
	}

	return time.Since(start)
}

func (s *SearchIndexBuilder) ExportAndSaveAsJSON() error {
	json, err := json.Marshal(s.searchIndex)

	destPath := filepath.Join(s.outputDirectory, s.generatedDirectory, s.fileName)
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	if err := os.WriteFile(destPath, json, 0644); err != nil {
		return fmt.Errorf("%s: %w", s.fileName, err)
	}

	if err != nil {
		return err
	}
	return nil
}
