package crawler

import (
	"gobsidian/internal/models"
	"regexp"
	"strings"
)

func (vc *VaultCrawler) populateExcalidrawImages(embeddedFiles map[string]string) []models.Image {
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
