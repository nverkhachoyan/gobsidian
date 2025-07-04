package parsers

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"gobsidian/internal/models"
	"gobsidian/internal/utils"

	"github.com/goccy/go-yaml"
)

type Parser interface {
	ParseNote(filePath string, fileName string) (ParsedNote, error)
}

type ObsidianParser struct {
	InputDirectory string
	Regexes        *models.ParserRegexes
}

func NewObsidianParser(c *ObsidianParser) *ObsidianParser {
	return &ObsidianParser{
		InputDirectory: c.InputDirectory,
		Regexes:        c.Regexes,
	}
}

type YamlFrontmatter struct {
	Title     string `yaml:"title"`
	Date      string `yaml:"date"`
	Author    string `yaml:"author"`
	UpdatedAt string `yaml:"updatedAt"`
}

type Frontmatter struct {
	title     string
	author    string
	date      time.Time
	updatedAt *time.Time
}

type ParsedNote struct {
	Title     string
	Author    string
	RawBody   []byte
	Date      time.Time
	UpdatedAt *time.Time
	Images    []models.Image
	Tags      []models.Tag
	Wikilinks []string
}

func (p *ObsidianParser) ParseNote(filePath string, fileName string) (ParsedNote, error) {
	markdownInput, err := os.ReadFile(filePath)
	if err != nil {
		return ParsedNote{}, err
	}

	frontmatterBytes, bodyBytes, err := p.splitFrontmatterAndBody(markdownInput)
	if err != nil {
		return ParsedNote{}, fmt.Errorf("failed to extract content from %s: %w", filePath, err)
	}

	frontmatter, err := p.parseFrontmatter(frontmatterBytes, filePath, fileName)
	if err != nil {
		return ParsedNote{}, fmt.Errorf("failed to parse frontmatter from %s: %w", filePath, err)
	}

	wikilinks := p.extractWikilinks(bodyBytes)
	images := p.extractObsidianImages(bodyBytes)
	tags := p.extractTags(bodyBytes)

	return ParsedNote{
		Title:     frontmatter.title,
		Author:    frontmatter.author,
		RawBody:   bodyBytes,
		Date:      frontmatter.date,
		UpdatedAt: frontmatter.updatedAt,
		Images:    images,
		Tags:      tags,
		Wikilinks: wikilinks,
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
	var frontmatter YamlFrontmatter

	if err := yaml.Unmarshal(frontmatterBytes, &frontmatter); err != nil {
		return Frontmatter{}, fmt.Errorf("failed to unmarshal frontmatter: %w", err)
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

	// Fallback to file mod time if parsing fails or date is not set
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
	}, nil
}
