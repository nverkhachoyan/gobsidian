package crawler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	lzstring "github.com/daku10/go-lz-string"
	"github.com/goccy/go-yaml"
)

func (vc *VaultCrawler) processMarkdownFile(node *models.VaultNode) {
	fileContent, err := os.ReadFile(node.Path)
	if err != nil {
		return
	}

	frontMatterBytes, bodyBytes := vc.splitFrontmatterAndBody(fileContent)

	node.Markdown = string(bodyBytes)
	bodyStr := string(bodyBytes)

	vc.parseAndSetFrontmatter(frontMatterBytes, node)

	node.IsLandingPageNote = slices.Contains(node.Tags, vc.landingPageTag)
	if slices.Contains(node.Tags, "excalidraw") {
		node.NoteType = models.NoteTypeExcalidraw
	}

	if node.NoteType != models.NoteTypeExcalidraw {
		vc.parseAndSetWikilinks(bodyStr, node)
		vc.parseAndSetMarkdownLinks(bodyStr, node)
		vc.parseAndSetTagsInMarkdownBody(bodyStr, node)

		// Word count (rough estimate)
		words := strings.Fields(bodyStr)
		node.WordCount = len(words)
	} else {
		vc.parseAndSetExcalidraw(bodyStr, node)
	}
}

func (vc *VaultCrawler) parseAndSetExcalidraw(bodyStr string, node *models.VaultNode) {
	textElements := parseTextElements(bodyStr)
	embeddedFilesMap := parseEmbeddedFiles(bodyStr)
	initialDataBase64 := extractExcalidrawBase64(bodyStr)
	_ = vc.populateExcalidrawImages(embeddedFilesMap)

	if initialDataBase64 == "" {
		fmt.Printf("no excalidraw base64 data found in %s\n", node.Path)
		fmt.Println("could not find excalidraw base64 data")
	}

	initialDataString, err := lzstring.DecompressFromBase64(initialDataBase64)
	if err != nil {
		fmt.Printf("failed to decompress excalidraw base64 data: %v\n", err)

	}

	var initialData models.InitialDataExcalidraw
	if err := json.Unmarshal([]byte(initialDataString), &initialData); err != nil {
		fmt.Printf("failed to unmarshal excalidraw initial data: %v\n", err)

	}

	node.Excalidraw = &models.Excalidraw{

		InitialData:   &initialData,
		TextElements:  textElements,
		EmbeddedFiles: embeddedFilesMap,
	}
}

func (vc *VaultCrawler) splitFrontmatterAndBody(markdownInput []byte) ([]byte, []byte) {
	frontmatterMatch := vc.regexes.FrontmatterRegex.FindSubmatch(markdownInput)

	if len(frontmatterMatch) > 1 {
		markdownInput = bytes.ReplaceAll(markdownInput, frontmatterMatch[0], []byte(""))
		return frontmatterMatch[1], markdownInput
	}

	return nil, markdownInput
}

func (vc *VaultCrawler) parseAndSetFrontmatter(frontmatterBytes []byte, node *models.VaultNode) {
	var warnings []string

	var genericFrontmatter map[string]any
	if err := yaml.Unmarshal(frontmatterBytes, &genericFrontmatter); err != nil {
		fmt.Printf("failed to unmarshal frontmatter: %v", err)
	}

	node.GenericFrontmatter = genericFrontmatter

	if title, ok := genericFrontmatter["title"].(string); ok {
		node.Title = title
	}
	if author, ok := genericFrontmatter["author"].(string); ok {
		node.Author = author
	}
	if date, ok := genericFrontmatter["date"].(string); ok {
		var postDate time.Time
		if date != "" {
			layouts := []string{"2006-01-02", time.RFC3339}
			for _, layout := range layouts {
				t, err := time.Parse(layout, date)
				if err == nil {
					postDate = t
					break
				}
			}
		}
		if postDate.IsZero() {
			fileInfo, err := os.Stat(node.Path)
			if err != nil {
				fmt.Printf("failed to get file info %s: %v", node.Path, err)
			}
			postDate = fileInfo.ModTime()
		}

		node.Date = postDate
	}

	if updatedAt, ok := genericFrontmatter["date"].(string); ok {
		var updatedAtDate time.Time
		if updatedAt != "" {
			layouts := []string{"2006-01-02", time.RFC3339}
			for _, layout := range layouts {
				t, err := time.Parse(layout, updatedAt)
				if err == nil {
					updatedAtDate = t
					break
				}
			}
		}

		if updatedAtDate.IsZero() {
			fileInfo, err := os.Stat(node.Path)
			if err != nil {
				fmt.Printf("failed to get file info %s: %v", node.Path, err)
			}
			updatedAtDate = fileInfo.ModTime()
		}

		node.UpdatedAt = updatedAtDate
	}

	cssClasses, cssWarnings := vc.parseCssClasses(genericFrontmatter)
	warnings = append(warnings, cssWarnings...)
	node.CssClasses = cssClasses

	var stringTags []string
	if tags, ok := genericFrontmatter["tags"].([]any); ok {
		for _, tag := range tags {
			if strTag, ok := tag.(string); ok {
				stringTags = append(stringTags, strTag)
			}
		}
	}
	node.Tags = append(node.Tags, stringTags...)
}

func (vc *VaultCrawler) parseAndSetWikilinks(bodyStr string, node *models.VaultNode) {
	wikilinks := vc.regexes.WikilinkRegex.FindAllStringSubmatch(bodyStr, -1)
	for i, match := range wikilinks {
		if len(match) > 1 {
			linkText := match[1]
			display := ""
			target := linkText

			// Handle display text: [[target|display]]
			if strings.Contains(linkText, "|") {
				parts := strings.SplitN(linkText, "|", 2)
				target = parts[0]
				display = parts[1]
			}

			link := models.Link{
				Target:  target,
				Display: display,
				Type:    "wikilink",
				Line:    vc.findLineNumber(bodyStr, match[0], i),
			}
			node.Links = append(node.Links, link)
		}
	}
}

func (vc *VaultCrawler) parseAndSetMarkdownLinks(bodyStr string, node *models.VaultNode) {
	mdlinks := vc.regexes.MarkdownRegex.FindAllStringSubmatch(bodyStr, -1)
	for i, match := range mdlinks {
		if len(match) > 2 {
			link := models.Link{
				Target:  match[2],
				Display: match[1],
				Type:    "markdown",
				Line:    vc.findLineNumber(bodyStr, match[0], i),
			}
			node.Links = append(node.Links, link)
		}
	}
}

func (vc *VaultCrawler) parseAndSetTagsInMarkdownBody(bodyStr string, node *models.VaultNode) {
	tags := vc.regexes.HashtagRegex.FindAllStringSubmatch(bodyStr, -1)
	for _, match := range tags {
		if len(match) > 1 && !slices.Contains(node.Tags, match[1]) {
			node.Tags = append(node.Tags, match[1])
		}
	}
}

func (vc *VaultCrawler) parseCssClasses(genericFrontmatter map[string]any) ([]string, []string) {
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

func (vc *VaultCrawler) resolveLinks() {
	for _, node := range vc.fileIndex {
		if node.IsDir {
			continue
		}

		for i := range node.Links {
			link := &node.Links[i]

			if link.Type == "wikilink" {
				resolvedNode := vc.resolveWikilink(node, link.Target)
				if resolvedNode != nil {
					link.IsResolved = true
					link.ResolvedPath = resolvedNode.Path
					link.TargetID = resolvedNode.ID
					link.TargetNode = resolvedNode
					link.URL = resolvedNode.URL

					backlink := models.Link{
						Target:       node.BaseName,
						TargetID:     node.ID,
						TargetNode:   node,
						Type:         "backlink",
						ResolvedPath: node.Path,
						URL:          node.URL,
						IsResolved:   true,
						Line:         0,
					}
					resolvedNode.Backlinks = append(resolvedNode.Backlinks, backlink)
				}
			}
		}
	}
}

func (vc *VaultCrawler) resolveWikilink(sourceNode *models.VaultNode, linkTarget string) *models.VaultNode {
	if strings.Contains(linkTarget, "/") || strings.Contains(linkTarget, "\\") {
		return vc.resolvePathBasedLink(linkTarget)
	}

	// For basename-only links, we follow Obsidian's search order
	candidates, exists := vc.nameIndex[linkTarget]
	if !exists {
		return nil
	}

	// If only one candidate exists, return it immediately
	if len(candidates) == 1 {
		return candidates[0]
	}

	// Multiple candidates - apply Obsidian's resolution priority
	sourceDir := filepath.Dir(sourceNode.Path)

	// 1. Check vault root first
	if resolved := vc.findInDirectory(candidates, vc.RootPath); resolved != nil {
		return resolved
	}

	// 2. Check source file's parent directory
	if resolved := vc.findInDirectory(candidates, sourceDir); resolved != nil {
		return resolved
	}

	// 3. Check subdirectories of source file's parent directory
	if resolved := vc.findInSubdirectories(candidates, sourceDir); resolved != nil {
		return resolved
	}

	// 4. Search everywhere else in the vault (breadth-first from root)
	if resolved := vc.findInVault(candidates, sourceDir); resolved != nil {
		return resolved
	}

	// If still not found, return the first candidate as fallback
	// (this shouldn't happen, but just in case)
	return candidates[0]
}

func (vc *VaultCrawler) resolvePathBasedLink(linkTarget string) *models.VaultNode {
	linkTarget = vc.normalizeDir(linkTarget)

	if node, exists := vc.depthFileIndex[linkTarget]; exists && !node.IsDir {
		return node
	}

	// Try with .md extension
	if !strings.HasSuffix(linkTarget, ".md") {
		if node, exists := vc.depthFileIndex[linkTarget+".md"]; exists && !node.IsDir {
			return node
		}
	}

	return nil
}
