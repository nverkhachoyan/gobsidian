package builders

import (
	"encoding/json"
	"fmt"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/charmbracelet/log"
)

type TagsBuilder struct {
	Logger             *log.Logger
	tags               []models.Tag
	outputDirectory    string
	generatedDirectory string
	fileName           string
}

func NewTagsBuilder(logger *log.Logger, outputDirectory string, generatedDirectory string, fileName string) *TagsBuilder {
	return &TagsBuilder{
		Logger:             logger,
		outputDirectory:    outputDirectory,
		generatedDirectory: generatedDirectory,
		fileName:           fileName,
	}
}

func (t *TagsBuilder) Build(fileIndex map[string]*models.VaultNode) time.Duration {
	start := time.Now()
	tagsMap := make(map[string]*models.Tag)
	for _, node := range fileIndex {
		for _, tag := range node.Tags {
			tag := &models.Tag{
				Name: tag,
				Slug: tag,
			}
			tagsMap[tag.Slug] = tag
			tagsMap[tag.Slug].Count++
		}
	}

	for _, tag := range tagsMap {
		t.tags = append(t.tags, *tag)
	}

	sort.Slice(t.tags, func(i, j int) bool {
		return t.tags[i].Name < t.tags[j].Name
	})

	return time.Since(start)
}

func (t *TagsBuilder) ExportAndSaveAsJSON() error {
	json, err := json.Marshal(t.tags)

	destPath := filepath.Join(t.outputDirectory, t.generatedDirectory, t.fileName)
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	if err := os.WriteFile(destPath, json, 0644); err != nil {
		return fmt.Errorf("%s: %w", t.fileName, err)
	}

	if err != nil {
		return err
	}
	return nil
}
