package parsers

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"gobsidian/internal/repository"
	"gobsidian/internal/utils"

	"github.com/goccy/go-yaml"
)

type Parser interface {
	ParseNote(filePath string, fileName string) (models.ParsedNote, error)
	ParseNotes(shallowNotes []*models.ScannedNote, notesRepository *repository.NoteRepository) time.Duration
}

type ObsidianParser struct {
	InputDirectory string
	Regexes        *config.Regexes
}

func NewObsidianParser(c *ObsidianParser) *ObsidianParser {
	return &ObsidianParser{
		InputDirectory: c.InputDirectory,
		Regexes:        c.Regexes,
	}
}

type ParseResult struct {
	Note *models.ParsedNote
	Err  error
}

type YamlFrontmatter struct {
	Title     string `yaml:"title"`
	Date      string `yaml:"date"`
	Author    string `yaml:"author"`
	UpdatedAt string `yaml:"updatedAt"`
}

type Frontmatter struct {
	title      string
	author     string
	date       time.Time
	updatedAt  *time.Time
	warnings   []string
	cssClasses []string
}

func (p *ObsidianParser) ParseNote(filePath string, fileName string) (models.ParsedNote, error) {
	markdownInput, err := os.ReadFile(filePath)
	if err != nil {
		return models.ParsedNote{}, err
	}

	frontmatterBytes, bodyBytes, err := p.splitFrontmatterAndBody(markdownInput)
	if err != nil {
		return models.ParsedNote{}, fmt.Errorf("failed to extract content from %s: %w", filePath, err)
	}

	frontmatter, err := p.parseFrontmatter(frontmatterBytes, filePath, fileName)
	if err != nil {
		return models.ParsedNote{}, fmt.Errorf("failed to parse frontmatter from %s: %w", filePath, err)
	}

	wikilinks := p.extractWikilinks(bodyBytes)
	images := p.extractObsidianImages(bodyBytes)
	tags := p.extractTags(bodyBytes)

	var warnings []string
	var missingFiles []string

	warnings = append(warnings, frontmatter.warnings...)

	for _, image := range images {
		if image.RelativePath != "" {
			imagePath := filepath.Join(p.InputDirectory, image.RelativePath)
			if _, err := os.Stat(imagePath); os.IsNotExist(err) {
				missingFiles = append(missingFiles, image.RelativePath)
			}
		}
	}

	return models.ParsedNote{
		Title:        frontmatter.title,
		Author:       frontmatter.author,
		RawBody:      bodyBytes,
		Date:         frontmatter.date,
		UpdatedAt:    frontmatter.updatedAt,
		Images:       images,
		Tags:         tags,
		Wikilinks:    wikilinks,
		Warnings:     warnings,
		MissingFiles: missingFiles,
		CssClasses:   frontmatter.cssClasses,
	}, nil
}

func (p *ObsidianParser) splitFrontmatterAndBody(markdownInput []byte) ([]byte, []byte, error) {
	frontmatterMatch := p.Regexes.FrontmatterRegex.FindSubmatch(markdownInput)

	if len(frontmatterMatch) > 1 {
		markdownInput = bytes.ReplaceAll(markdownInput, frontmatterMatch[0], []byte(""))
		return frontmatterMatch[1], markdownInput, nil
	}

	return nil, markdownInput, nil
}

func (p *ObsidianParser) extractWikilinks(markdownInput []byte) []string {
	var linkedRelativePaths []string
	wikilinkMatches := p.Regexes.WikilinkRegex.FindAllSubmatch(markdownInput, -1)
	for _, match := range wikilinkMatches {
		linkedRelativePaths = append(linkedRelativePaths, string(match[1]))
	}

	return linkedRelativePaths
}

func (p *ObsidianParser) extractObsidianImages(markdownInput []byte) []models.Image {
	var images []models.Image
	imageMatches := p.Regexes.ObsidianImageRegex.FindAllSubmatch(markdownInput, -1)
	for _, match := range imageMatches {
		var image models.Image
		if len(match) > 1 {
			image.RelativePath = string(match[1])
			image.Alt = string(match[2])

		}
		if len(match) > 2 {
			image.Width = string(match[2])
		}
		if len(match) > 3 {
			image.Height = string(match[3])
		}

		images = append(images, image)
	}
	return images
}

func (p *ObsidianParser) extractTags(markdownInput []byte) []models.Tag {
	var tags []models.Tag
	tagMatches := p.Regexes.HashtagRegex.FindAllSubmatch(markdownInput, -1)
	for _, match := range tagMatches {
		tagName := string(match[1])
		tags = append(tags, models.Tag{
			Name: tagName,
			Slug: utils.Slugify(tagName),
		})
	}

	return tags
}

func (p *ObsidianParser) parseFrontmatter(frontmatterBytes []byte, filePath string, fileName string) (Frontmatter, error) {
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

	cssClasses, cssWarnings := p.extractCssClasses(genericFrontmatter)
	warnings = append(warnings, cssWarnings...)

	return Frontmatter{
		author:     frontmatter.Author,
		title:      title,
		date:       postDate,
		updatedAt:  updatedAt,
		cssClasses: cssClasses,
		warnings:   warnings,
	}, nil
}

func (p *ObsidianParser) extractCssClasses(genericFrontmatter map[string]any) ([]string, []string) {
	var warnings []string
	var cssClasses []string

	rawClasses, ok := genericFrontmatter["cssclasses"]
	if !ok {
		return cssClasses, warnings
	}

	switch v := rawClasses.(type) {
	case string:
		cssClasses = append(cssClasses, v)
	case []interface{}:
		for _, item := range v {
			if classStr, ok := item.(string); ok {
				cssClasses = append(cssClasses, classStr)
			} else {
				warnings = append(warnings, fmt.Sprintf("cssclasses property contains a non-string item: %v", item))
			}
		}
	case []string:
		cssClasses = v
	default:
		warnings = append(warnings, fmt.Sprintf("cssclasses property has an unexpected type: %T", v))
	}

	return cssClasses, warnings
}

func (p *ObsidianParser) ParseNotes(
	shallowNotes []*models.ScannedNote,
	notesRepository *repository.NoteRepository,
) time.Duration {
	start := time.Now()
	numWorkers := min(runtime.NumCPU(), len(shallowNotes))
	notesChan := make(chan *models.ScannedNote, len(shallowNotes))
	resultsChan := make(chan ParseResult, len(shallowNotes))

	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for note := range notesChan {
				result := p.parseNoteWorker(note)
				resultsChan <- result
			}
		}()
	}

	for _, note := range shallowNotes {
		notesChan <- note
	}
	close(notesChan)

	wg.Wait()
	close(resultsChan)

	for result := range resultsChan {
		if result.Err != nil {
			continue
		}

		notesRepository.AddNote(result.Note)
	}

	return time.Since(start)
}

func (p *ObsidianParser) parseNoteWorker(
	scannedNote *models.ScannedNote,
) ParseResult {
	if scannedNote == nil {
		return ParseResult{
			Note: nil,
			Err:  fmt.Errorf("received nil scannedNote"),
		}
	}

	parsedNoteFromParser, err := p.ParseNote(scannedNote.FullPath, scannedNote.FileName)

	// passing as pointer is safe since each parseNoteWorker gets a different note
	parsedNote := &models.ParsedNote{
		ScannedNote: scannedNote,
	}

	if err != nil {
		return ParseResult{
			Note: parsedNote,
			Err:  err,
		}
	}

	if scannedNote.IsLandingPage {
		parsedNote.URL = "/"
	} else {
		URL := strings.TrimPrefix(scannedNote.FullPath, p.InputDirectory)
		URL = strings.TrimPrefix(URL, "/")
		URL = strings.TrimSuffix(URL, ".md")

		slugifiedURL := "/" + utils.Slugify(URL) + ".html"
		parsedNote.URL = slugifiedURL
	}

	parsedNote.Title = parsedNoteFromParser.Title
	parsedNote.Author = parsedNoteFromParser.Author
	parsedNote.RawBody = parsedNoteFromParser.RawBody
	parsedNote.Date = parsedNoteFromParser.Date
	parsedNote.UpdatedAt = parsedNoteFromParser.UpdatedAt
	parsedNote.Tags = parsedNoteFromParser.Tags
	parsedNote.Images = parsedNoteFromParser.Images
	parsedNote.Wikilinks = parsedNoteFromParser.Wikilinks
	parsedNote.Warnings = parsedNoteFromParser.Warnings
	parsedNote.MissingFiles = parsedNoteFromParser.MissingFiles
	parsedNote.CssClasses = parsedNoteFromParser.CssClasses
	return ParseResult{
		Note: parsedNote,
		Err:  nil,
	}
}

type ParseError struct {
	FilePath string
	Message  string
	Cause    error
}

func (e ParseError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("parse error in %s: %s (caused by: %v)", e.FilePath, e.Message, e.Cause)
	}
	return fmt.Sprintf("parse error in %s: %s", e.FilePath, e.Message)
}

func (e ParseError) Unwrap() error {
	return e.Cause
}
