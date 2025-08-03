package crawler

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/goccy/go-yaml"
)

func (vc *VaultCrawler) processMarkdownFile(node *VaultNode) {
	fileContent, err := os.ReadFile(node.Path)
	if err != nil {
		return
	}

	frontMatterBytes, bodyBytes, err := vc.splitFrontmatterAndBody(fileContent)
	if err != nil {
		fmt.Println("failed to split frontmatter and body")
	}

	node.Markdown = string(bodyBytes)
	bodyStr := string(bodyBytes)

	frontMatter, err := vc.parseFrontmatter(frontMatterBytes, node.Path, node.BaseName)
	if err != nil {
		fmt.Println("failed when parsing frontmatter", err)
	}

	node.Title = frontMatter.title
	node.Date = frontMatter.date

	// Extract wikilinks
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

			link := Link{
				Target:  target,
				Display: display,
				Type:    "wikilink",
				Line:    vc.findLineNumber(bodyStr, match[0], i),
			}
			node.Links = append(node.Links, link)
		}
	}

	// Extract markdown links
	mdlinks := vc.regexes.WikilinkRegex.FindAllStringSubmatch(bodyStr, -1)
	for i, match := range mdlinks {
		if len(match) > 2 {
			link := Link{
				Target:  match[2],
				Display: match[1],
				Type:    "markdown",
				Line:    vc.findLineNumber(bodyStr, match[0], i),
			}
			node.Links = append(node.Links, link)
		}
	}

	// Extract tags
	// TODO: make sure to get tags from frontmatter as well
	tags := vc.regexes.HashtagRegex.FindAllStringSubmatch(bodyStr, -1)
	for _, match := range tags {
		if len(match) > 1 {
			node.Tags = append(node.Tags, match[1])
		}
	}

	// Calculate word count (rough estimate)
	words := strings.Fields(bodyStr)
	node.WordCount = len(words)
}

func (vc *VaultCrawler) splitFrontmatterAndBody(markdownInput []byte) ([]byte, []byte, error) {
	frontmatterMatch := vc.regexes.FrontmatterRegex.FindSubmatch(markdownInput)

	if len(frontmatterMatch) > 1 {
		markdownInput = bytes.ReplaceAll(markdownInput, frontmatterMatch[0], []byte(""))
		return frontmatterMatch[1], markdownInput, nil
	}

	return nil, markdownInput, nil
}

func (p *VaultCrawler) parseFrontmatter(frontmatterBytes []byte, filePath string, fileName string) (Frontmatter, error) {
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

func (vc *VaultCrawler) extractCssClasses(genericFrontmatter map[string]any) ([]string, []string) {
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
	for _, node := range vc.FileIndex {
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

					backlink := Link{
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

// resolveWikilink implements Obsidian's wikilink resolution algorithm
func (vc *VaultCrawler) resolveWikilink(sourceNode *VaultNode, linkTarget string) *VaultNode {
	// Handle path-based links (contains / or \)
	if strings.Contains(linkTarget, "/") || strings.Contains(linkTarget, "\\") {
		// TODO: we need an index of all possible nested links
		return vc.resolvePathBasedLink(sourceNode, linkTarget)
	}

	// For basename-only links, follow Obsidian's search order
	candidates, exists := vc.NameIndex[linkTarget]
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
	// (this shouldn't happen in normal Obsidian behavior, but provides robustness)
	return candidates[0]
}

// resolvePathBasedLink handles links that contain path separators
func (vc *VaultCrawler) resolvePathBasedLink(sourceNode *VaultNode, linkTarget string) *VaultNode {
	linkTarget = vc.normalizeDir(linkTarget)

	// Try absolute path from vault root
	absolutePath := filepath.Join(vc.RootPath, linkTarget)
	if node, exists := vc.FileIndex[absolutePath]; exists && !node.IsDir {
		return node
	}

	// Try with .md extension if not present
	if !strings.HasSuffix(linkTarget, ".md") {
		absolutePathMd := absolutePath + ".md"
		if node, exists := vc.FileIndex[absolutePathMd]; exists && !node.IsDir {
			return node
		}
	}

	// Try relative path from source file's directory
	sourceDir := filepath.Dir(sourceNode.Path)
	relativePath := filepath.Join(sourceDir, linkTarget)
	if node, exists := vc.FileIndex[relativePath]; exists && !node.IsDir {
		return node
	}

	// Try relative path with .md extension
	if !strings.HasSuffix(linkTarget, ".md") {
		relativePathMd := relativePath + ".md"
		if node, exists := vc.FileIndex[relativePathMd]; exists && !node.IsDir {
			return node
		}
	}

	// Try relative path from source file's directory WITH DEPTH INDEX
	if node, exists := vc.DepthFileIndex[linkTarget]; exists && !node.IsDir {
		return node
	}

	// Try relative path with .md extension
	if !strings.HasSuffix(linkTarget, ".md") {
		if node, exists := vc.DepthFileIndex[linkTarget+".md"]; exists && !node.IsDir {
			return node
		}
	}

	return nil
}

// findInDirectory looks for candidates in a specific directory
func (vc *VaultCrawler) findInDirectory(candidates []*VaultNode, dirPath string) *VaultNode {
	for _, candidate := range candidates {
		if filepath.Dir(candidate.Path) == dirPath {
			return candidate
		}
	}
	return nil
}

// findInSubdirectories looks for candidates in subdirectories of the given directory
func (vc *VaultCrawler) findInSubdirectories(candidates []*VaultNode, parentDir string) *VaultNode {
	for _, candidate := range candidates {
		candidateDir := filepath.Dir(candidate.Path)

		// Check if candidate is in a subdirectory of parentDir
		if vc.isSubdirectory(candidateDir, parentDir) {
			return candidate
		}
	}
	return nil
}

// findInVault searches the entire vault, excluding already searched locations
func (vc *VaultCrawler) findInVault(candidates []*VaultNode, excludeDir string) *VaultNode {
	// Sort candidates by depth (prefer files closer to root)
	sortedCandidates := make([]*VaultNode, len(candidates))
	copy(sortedCandidates, candidates)

	// Simple depth-based sorting (count path separators)
	for i := 0; i < len(sortedCandidates)-1; i++ {
		for j := i + 1; j < len(sortedCandidates); j++ {
			depthI := strings.Count(sortedCandidates[i].Path, string(filepath.Separator))
			depthJ := strings.Count(sortedCandidates[j].Path, string(filepath.Separator))

			if depthI > depthJ {
				sortedCandidates[i], sortedCandidates[j] = sortedCandidates[j], sortedCandidates[i]
			}
		}
	}

	for _, candidate := range sortedCandidates {
		candidateDir := filepath.Dir(candidate.Path)

		// Skip if we already checked this location
		if candidateDir == vc.RootPath ||
			candidateDir == excludeDir ||
			vc.isSubdirectory(candidateDir, excludeDir) {
			continue
		}

		return candidate
	}

	return nil
}

func (vc *VaultCrawler) findLineNumber(content, needle string, occurrence int) int {
	lines := strings.Split(content, "\n")
	found := 0

	for i, line := range lines {
		if strings.Contains(line, needle) {
			if found == occurrence {
				return i + 1
			}
			found++
		}
	}
	return 0
}
