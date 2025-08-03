package parsers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"os"
	"regexp"
	"strings"
	"time"

	lzstring "github.com/daku10/go-lz-string"
	"github.com/goccy/go-yaml"
)

type ExcalidrawParser struct {
	InputDirectory  string
	OutputDirectory string
	Regexes         *config.Regexes
	BaseURL         string
	Port            int
}

func NewExcalidrawParser(
	inputDirectory string,
	outputDirectory string,
	baseURL string,
	port int,
	regexes *config.Regexes,
) *ExcalidrawParser {
	return &ExcalidrawParser{
		InputDirectory:  inputDirectory,
		OutputDirectory: outputDirectory,
		BaseURL:         baseURL,
		Port:            port,
		Regexes:         regexes,
	}
}

func (p *ExcalidrawParser) Parse(filePath string, filename string) (models.ParsedNote, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return models.ParsedNote{}, err
	}

	frontmatterBytes, bodyBytes, err := p.SplitFrontmatterAndBody(file)
	if err != nil {
		return models.ParsedNote{}, fmt.Errorf("failed to split frontmatter and body from %s: %w", filePath, err)
	}

	frontmatter, err := p.parseFrontmatter(frontmatterBytes, filePath, filename)
	if err != nil {
		return models.ParsedNote{}, fmt.Errorf("failed to parse frontmatter from %s: %w", filePath, err)
	}

	markdown := string(bodyBytes)

	// Extract text elements and embedded files
	textElements := parseTextElements(markdown)
	embeddedFilesMap := parseEmbeddedFiles(markdown)
	initialDataBase64 := extractExcalidrawBase64(markdown)
	images := p.populateExcalidrawImages(embeddedFilesMap)

	if initialDataBase64 == "" {
		fmt.Printf("no excalidraw base64 data found in %s\n", filePath)
		return models.ParsedNote{}, fmt.Errorf("could not find excalidraw base64 data")
	}

	// Decompress the excalidraw base64 data
	initialDataString, err := lzstring.DecompressFromBase64(initialDataBase64)
	if err != nil {
		fmt.Printf("failed to decompress excalidraw base64 data: %v\n", err)
		return models.ParsedNote{}, fmt.Errorf("failed to decompress excalidraw base64 data: %w", err)
	}

	var initialData models.InitialDataExcalidraw
	if err := json.Unmarshal([]byte(initialDataString), &initialData); err != nil {
		fmt.Printf("failed to unmarshal excalidraw initial data: %v\n", err)
		return models.ParsedNote{}, fmt.Errorf("failed to unmarshal excalidraw initial data: %w", err)
	}

	return models.ParsedNote{
		Title:     frontmatter.title,
		Date:      frontmatter.date,
		Author:    frontmatter.author,
		UpdatedAt: frontmatter.updatedAt,
		Excalidraw: &models.Excalidraw{
			InitialData:   &initialData,
			TextElements:  textElements,
			EmbeddedFiles: embeddedFilesMap,
		},
		Images: images,
	}, nil
}

func (p *ExcalidrawParser) SplitFrontmatterAndBody(markdownInput []byte) ([]byte, []byte, error) {
	frontmatterMatch := p.Regexes.FrontmatterRegex.FindSubmatch(markdownInput)

	if len(frontmatterMatch) > 1 {
		markdownInput = bytes.ReplaceAll(markdownInput, frontmatterMatch[0], []byte(""))
		return frontmatterMatch[1], markdownInput, nil
	}

	return nil, markdownInput, nil
}

func (p *ExcalidrawParser) parseFrontmatter(frontmatterBytes []byte, filePath string, fileName string) (Frontmatter, error) {
	var warnings []string

	var genericFrontmatter map[string]any
	if err := yaml.Unmarshal(frontmatterBytes, &genericFrontmatter); err != nil {
		return Frontmatter{}, fmt.Errorf("failed to unmarshal frontmatter: %w", err)
	}

	var frontmatter YamlFrontmatter
	if err := yaml.Unmarshal(frontmatterBytes, &frontmatter); err != nil {
		if title, ok := genericFrontmatter["title"].(string); ok {
			frontmatter.Title = title
		}
		if author, ok := genericFrontmatter["author"].(string); ok {
			frontmatter.Author = author
		}
		if date, ok := genericFrontmatter["date"].(string); ok {
			frontmatter.Date = date
		}
		if updatedAt, ok := genericFrontmatter["updatedAt"].(string); ok {
			frontmatter.UpdatedAt = updatedAt
		}

		warnings = append(warnings, fmt.Sprintf("frontmatter parsing issues detected, some fields may be ignored: %v", err))
	}

	knownFields := map[string]bool{
		"title":      true,
		"date":       true,
		"author":     true,
		"updatedAt":  true,
		"tags":       true,
		"cssclasses": true,
		"aliases":    true,
	}

	for field := range genericFrontmatter {
		if !knownFields[field] {
			warnings = append(warnings, fmt.Sprintf("unknown frontmatter field: %s", field))
		}
	}

	title := frontmatter.Title
	if title == "" {
		title = fileName
	}

	var postDate time.Time
	if frontmatter.Date != "" {
		layouts := []string{"2006-01-02", time.RFC3339}
		for _, layout := range layouts {
			t, err := time.Parse(layout, frontmatter.Date)
			if err == nil {
				postDate = t
				break
			}
		}
	}

	if postDate.IsZero() {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return Frontmatter{}, fmt.Errorf("failed to get file info %s: %w", filePath, err)
		}
		postDate = fileInfo.ModTime()
	}

	var updatedAt *time.Time
	if frontmatter.UpdatedAt != "" {
		layouts := []string{"2006-01-02", time.RFC3339}
		for _, layout := range layouts {
			t, err := time.Parse(layout, frontmatter.UpdatedAt)
			if err == nil {
				updatedAt = &t
				break
			}
		}
	}

	return Frontmatter{
		author:    frontmatter.Author,
		title:     title,
		date:      postDate,
		updatedAt: updatedAt,

		warnings: warnings,
	}, nil
}

func (p *ExcalidrawParser) populateExcalidrawImages(embeddedFiles map[string]string) []models.Image {
	var images []models.Image

	for _, file := range embeddedFiles {
		image := models.Image{
			RelativePath: file,
		}
		images = append(images, image)
	}

	return images
}

func parseTextElements(content string) map[string]string {
	re := regexp.MustCompile(`(?m)^(.+?)\s*\^([a-zA-Z0-9]+)$`)
	sectionRe := regexp.MustCompile(`(?s)## Text Elements\n(.*?)\n(%%|##)`)

	elements := make(map[string]string)
	sectionMatch := sectionRe.FindStringSubmatch(content)
	if len(sectionMatch) < 2 {
		return elements
	}

	matches := re.FindAllStringSubmatch(sectionMatch[1], -1)
	for _, match := range matches {
		elements[match[2]] = match[1]
	}
	return elements
}

func parseEmbeddedFiles(content string) map[string]string {
	re := regexp.MustCompile(`(?m)^([a-zA-Z0-9]+):\s*\[\[(.+?)\]\]$`)
	sectionRe := regexp.MustCompile(`(?s)## Embedded Files\n(.*?)\n(%%|##)`)

	files := make(map[string]string)
	sectionMatch := sectionRe.FindStringSubmatch(content)
	if len(sectionMatch) < 2 {
		return files
	}

	matches := re.FindAllStringSubmatch(sectionMatch[1], -1)
	for _, match := range matches {
		files[match[1]] = match[2]
	}
	return files
}

func extractExcalidrawBase64(content string) string {
	re := regexp.MustCompile("(?s)```compressed-json\n(.*?)\n```")
	match := re.FindStringSubmatch(content)
	if len(match) > 1 {
		parts := strings.Split(match[1], "\n\n")
		return strings.Join(parts, "")
	}
	return ""
}
