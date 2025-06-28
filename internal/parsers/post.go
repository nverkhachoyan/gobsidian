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

	"github.com/goccy/go-yaml"
)

type Parser interface {
	ParseBlogPost(filePath string) (models.BlogPost, []string, []string, error)
}

type MarkdownParser struct {
	Config config.Config
}

func NewMarkdownParser(c config.Config) *MarkdownParser {
	return &MarkdownParser{
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

func (p *MarkdownParser) ParseBlogPost(filePath string) (models.BlogPost, []string, []string, error) {
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

	linkedPosts := p.extractWikilinks(bodyWithoutFrontmatter)
	images := p.extractObsidianImages(bodyWithoutFrontmatter)
	tags := p.extractTags(bodyWithoutFrontmatter)

	// Generate a clean filename for the HTML output
	htmlFileName := utils.Slugify(processedFrontmatter.title) + ".html"
	if htmlFileName == ".html" { // Avoid empty file names if title is empty
		htmlFileName = "post-" + time.Now().Format("20060102150405") + ".html"
	}

	return models.BlogPost{
		Title:     processedFrontmatter.title,
		FileName:  htmlFileName,
		RawBody:   bodyWithoutFrontmatter,
		Date:      processedFrontmatter.date,
		Author:    frontmatter.Author,
		UpdatedAt: processedFrontmatter.updatedAt,
		Tags:      tags,
	}, images, linkedPosts, nil
}

func (p *MarkdownParser) processFrontmatter(frontmatter Frontmatter, filePath string) (processedFrontmatter, error) {
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

func (p *MarkdownParser) splitFrontmatterAndBody(markdownInput []byte) (Frontmatter, []byte, error) {
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

func (p *MarkdownParser) extractWikilinks(markdownInput []byte) []string {
	var linkedPosts []string
	wikilinkMatches := p.Config.RegexpConfig.WikilinkRegex.FindAllSubmatch(markdownInput, -1)
	for _, match := range wikilinkMatches {
		linkedPosts = append(linkedPosts, string(match[1]))
	}

	return linkedPosts
}

func (p *MarkdownParser) extractObsidianImages(markdownInput []byte) []string {
	var images []string
	imageMatches := p.Config.RegexpConfig.ObsidianImageRegex.FindAllSubmatch(markdownInput, -1)
	for _, match := range imageMatches {
		if len(match) > 1 {
			imageName := string(match[1])
			images = append(images, imageName)
		}
	}
	return images
}

func (p *MarkdownParser) extractTags(markdownInput []byte) []models.Tag {
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
