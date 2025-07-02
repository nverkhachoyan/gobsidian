package parsers

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"gobsidian/internal/utils"

	"sync/atomic"

	"github.com/goccy/go-yaml"
)

type Parser interface {
	ParseNote(filePath string, fileName string) (models.BlogPost, []models.Image, []string, error)
}

type ObsidianParser struct {
	Config config.Config
}

func NewObsidianParser(c config.Config) *ObsidianParser {
	return &ObsidianParser{
		Config: c,
	}
}

type Frontmatter struct {
	Title     string `yaml:"title"`
	Date      string `yaml:"date"`
	Author    string `yaml:"author"`
	UpdatedAt string `yaml:"updatedAt"`
}

type processedFrontmatter struct {
	title     string
	date      time.Time
	updatedAt *time.Time
}

var atomicIDcounter int64

func getNextID() int64 {
	return atomic.AddInt64(&atomicIDcounter, 1)
}

func (p *ObsidianParser) ParseNote(filePath string, fileName string) (models.BlogPost, []models.Image, []string, error) {
	markdownInput, err := os.ReadFile(filePath)
	if err != nil {
		return models.BlogPost{}, nil, nil, err
	}

	frontmatter, bodyWithoutFrontmatter, err := p.splitFrontmatterAndBody(markdownInput)
	if err != nil {
		return models.BlogPost{}, nil, nil, fmt.Errorf("failed to extract content from %s: %w", filePath, err)
	}

	processedFrontmatter, err := p.processFrontmatter(frontmatter, filePath)
	if err != nil {
		return models.BlogPost{}, nil, nil, fmt.Errorf("failed to process frontendmatter from %s: %w", filePath, err)
	}

	linkedRelativePaths := p.extractWikilinks(bodyWithoutFrontmatter)
	images := p.extractObsidianImages(bodyWithoutFrontmatter)
	tags := p.extractTags(bodyWithoutFrontmatter)

	htmlFileName := utils.Slugify(strings.TrimSuffix(fileName, ".md")) + ".html"

	inputDir := strings.TrimPrefix(p.Config.InputDirectory, "./")
	relativePath := strings.TrimPrefix(filePath, inputDir)
	relativePathWithoutExtension := strings.TrimPrefix(strings.TrimSuffix(relativePath, ".md"), "/")

	relativePathWithoutName := strings.TrimRight(relativePath, fileName)
	relativePathWithoutName = strings.TrimRight(relativePathWithoutName, "/")
	relativePathWithoutName = utils.Slugify(relativePathWithoutName)

	return models.BlogPost{
		ID:                      getNextID(),
		Title:                   processedFrontmatter.title,
		FileName:                htmlFileName,
		RawFileName:             strings.TrimSuffix(fileName, ".md"),
		RawBody:                 bodyWithoutFrontmatter,
		Date:                    processedFrontmatter.date,
		Author:                  frontmatter.Author,
		UpdatedAt:               processedFrontmatter.updatedAt,
		Tags:                    tags,
		RelativePath:            relativePathWithoutExtension,
		RelativePathWithoutName: relativePathWithoutName,
	}, images, linkedRelativePaths, nil
}

func (p *ObsidianParser) processFrontmatter(frontmatter Frontmatter, filePath string) (processedFrontmatter, error) {
	title := frontmatter.Title
	if title == "" {
		title = strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
		title = strings.ReplaceAll(title, "-", " ")
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
			return processedFrontmatter{}, fmt.Errorf("failed to get file info %s: %w", filePath, err)
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

	return processedFrontmatter{
		title:     title,
		date:      postDate,
		updatedAt: updatedAt,
	}, nil
}

func (p *ObsidianParser) splitFrontmatterAndBody(markdownInput []byte) (Frontmatter, []byte, error) {
	var frontmatter Frontmatter

	frontmatterMatch := p.Config.RegexpConfig.FrontmatterRegex.FindSubmatch(markdownInput)
	if len(frontmatterMatch) > 1 {
		if err := yaml.Unmarshal(frontmatterMatch[1], &frontmatter); err == nil {
			markdownInput = bytes.ReplaceAll(markdownInput, frontmatterMatch[0], []byte(""))
			markdownInput = []byte(strings.TrimSpace(string(markdownInput)))

			return frontmatter, markdownInput, nil
		}
	}

	return Frontmatter{}, markdownInput, nil
}

func (p *ObsidianParser) extractWikilinks(markdownInput []byte) []string {
	var linkedRelativePaths []string
	wikilinkMatches := p.Config.RegexpConfig.WikilinkRegex.FindAllSubmatch(markdownInput, -1)
	for _, match := range wikilinkMatches {
		linkedRelativePaths = append(linkedRelativePaths, string(match[1]))
	}

	return linkedRelativePaths
}

func (p *ObsidianParser) extractObsidianImages(markdownInput []byte) []models.Image {
	var images []models.Image
	imageMatches := p.Config.RegexpConfig.ObsidianImageRegex.FindAllSubmatch(markdownInput, -1)
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
	tagMatches := p.Config.RegexpConfig.HashtagRegex.FindAllSubmatch(markdownInput, -1)
	for _, match := range tagMatches {
		tagName := string(match[1])
		tags = append(tags, models.Tag{
			Name: tagName,
			Slug: utils.Slugify(tagName),
		})
	}

	return tags
}
