package builders

import (
	"encoding/json"
	"fmt"
	"gobsidian/internal/models"
	"gobsidian/internal/utils"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

type FileTreeBuilder struct {
	Logger *log.Logger
}

type FileTreeBuilderResult struct {
	Duration        time.Duration
	Root            *models.Folder
	NumberOfFolders int
	JSON            []byte
}

func NewFileTreeBuilder(logger *log.Logger) *FileTreeBuilder {
	return &FileTreeBuilder{Logger: logger}
}

func (ft *FileTreeBuilder) Build(
	notesByPath map[string]*models.ParsedNote,
) (*FileTreeBuilderResult, error) {
	start := time.Now()
	root := &models.Folder{
		Name:     "Home",
		Path:     "",
		Posts:    []*models.ParsedNote{},
		Children: make(map[string]*models.Folder),
	}
	numberOfFolders := 0

	for _, note := range notesByPath {
		if !note.IsInsideFolder {
			root.Posts = append(root.Posts, note)
			continue
		}

		htmlFileName := utils.Slugify(note.FileName) + ".html"
		url := strings.TrimRight(note.URL, htmlFileName)
		parts := strings.Split(url, string(filepath.Separator))
		currentNode := root

		for _, part := range parts {
			if part == "" {
				continue
			}

			lowercasePart := strings.ToLower(part)
			name := utils.Deslugify(lowercasePart)

			if _, ok := currentNode.Children[lowercasePart]; !ok {
				childPath := filepath.ToSlash(filepath.Join(currentNode.Path, lowercasePart))

				currentNode.Children[lowercasePart] = &models.Folder{
					Name:     name,
					Path:     childPath,
					Posts:    []*models.ParsedNote{},
					Children: make(map[string]*models.Folder),
				}
				numberOfFolders++
			}
			currentNode = currentNode.Children[lowercasePart]
		}

		currentNode.Posts = append(currentNode.Posts, note)
	}

	ft.sortFolderTree(root)

	jsonTree := ft.convertFolderToNode(root)
	jsonBytes, err := json.Marshal(jsonTree)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal folder tree to JSON: %w", err)
	}

	return &FileTreeBuilderResult{
		Duration:        time.Since(start),
		Root:            root,
		NumberOfFolders: numberOfFolders,
		JSON:            jsonBytes,
	}, nil
}

func (ft *FileTreeBuilder) sortFolderTree(folder *models.Folder) {
	sort.Slice(folder.Posts, func(i, j int) bool {
		return folder.Posts[i].Title < folder.Posts[j].Title
	})

	for _, child := range folder.Children {
		ft.sortFolderTree(child)
	}
}

func (ft *FileTreeBuilder) convertFolderToNode(folder *models.Folder) *models.FolderNode {
	node := &models.FolderNode{
		Name:     folder.Name,
		Path:     folder.Path,
		Children: make([]*models.FolderNode, 0),
		Files:    make([]*models.File, 0),
	}

	for _, post := range folder.Posts {
		node.Files = append(node.Files, &models.File{
			Name: post.Title,
			URL:  post.URL,
		})
	}
	sort.Slice(node.Files, func(i, j int) bool {
		return node.Files[i].Name < node.Files[j].Name
	})

	childKeys := make([]string, 0, len(folder.Children))
	for key := range folder.Children {
		childKeys = append(childKeys, key)
	}
	sort.Strings(childKeys)

	for _, key := range childKeys {
		childFolder := folder.Children[key]
		node.Children = append(node.Children, ft.convertFolderToNode(childFolder))
	}

	return node
}
