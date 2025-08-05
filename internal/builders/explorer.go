package builders

import (
	"encoding/json"
	"gobsidian/internal/models"
	"sort"
	"time"

	"github.com/charmbracelet/log"
)

type ExplorerBuilder struct {
	Logger      *log.Logger
	folderIndex map[int64]*models.FolderNode
}

type ExplorerBuilderResult struct {
	Duration        time.Duration
	Root            *models.FolderNode
	NumberOfFolders int
	JSON            []byte
}

func NewExplorerBuilder(logger *log.Logger) *ExplorerBuilder {
	return &ExplorerBuilder{Logger: logger,
		folderIndex: make(map[int64]*models.FolderNode, 0)}
}

func (ft *ExplorerBuilder) Build(rootNode *models.VaultNode) *ExplorerBuilderResult {
	start := time.Now()
	rootFolderNode := &models.FolderNode{}
	ft.buildFolder(rootNode, rootFolderNode)

	ft.sortFolderTree(rootFolderNode)

	filetreeJSON, err := json.Marshal(rootFolderNode)
	if err != nil {
		ft.Logger.Debug("Failed to unmarshal file tree", err)
	}

	return &ExplorerBuilderResult{
		Duration: time.Since(start),
		Root:     rootFolderNode,
		JSON:     filetreeJSON,
	}
}

func (ft *ExplorerBuilder) buildFolder(rootNode *models.VaultNode, rootFolderNode *models.FolderNode) {
	rootFolderNode.Name = rootNode.Name
	rootFolderNode.Path = rootNode.URL
	for _, node := range rootNode.Children {
		if node.IsDir {
			folder := &models.FolderNode{
				Name: node.Name,
				Path: node.URL,
			}
			rootFolderNode.Children = append(rootFolderNode.Children, folder)
			ft.folderIndex[node.ID] = folder
			ft.buildFolder(node, folder)
		} else if node.NoteType == models.NoteTypeMarkdown || node.NoteType == models.NoteTypeExcalidraw {
			rootFolderNode.Files = append(rootFolderNode.Files, &models.File{
				Name: node.BaseName,
				URL:  node.URL,
			})
		}
	}
}

func (ft *ExplorerBuilder) GetFolderIndex() map[int64]*models.FolderNode {
	return ft.folderIndex
}

func (ft *ExplorerBuilder) sortFolderTree(folder *models.FolderNode) {
	sort.Slice(folder.Files, func(i, j int) bool {
		return folder.Files[i].Name < folder.Files[j].Name
	})

	for _, child := range folder.Children {
		ft.sortFolderTree(child)
	}
}
