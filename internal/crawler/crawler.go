package crawler

import (
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"gobsidian/internal/utils"
	"mime"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type GobsidianFrontmatter struct {
	title      string
	author     string
	date       time.Time
	updatedAt  *time.Time
	warnings   []string
	cssClasses []string
	extra      map[string]any
}

type GenericFrontmatter map[string]any

type YamlFrontmatter struct {
	Title     string `yaml:"title"`
	Date      string `yaml:"date"`
	Author    string `yaml:"author"`
	UpdatedAt string `yaml:"updatedAt"`
	Extra     map[string]any
}

type VaultCrawler struct {
	RootPath       string
	landingPageTag string
	fileIndex      map[string]*models.VaultNode
	depthFileIndex map[string]*models.VaultNode
	nameIndex      map[string][]*models.VaultNode
	idToNodeIndex  map[int64]*models.VaultNode
	tagIndex       map[string][]*models.VaultNode
	regexes        *config.Regexes
}

func NewVaultCrawler(rootPath, landingPageTag string, regexes *config.Regexes) *VaultCrawler {
	return &VaultCrawler{
		RootPath:       rootPath,
		landingPageTag: landingPageTag,
		fileIndex:      make(map[string]*models.VaultNode),
		depthFileIndex: make(map[string]*models.VaultNode),
		nameIndex:      make(map[string][]*models.VaultNode),
		idToNodeIndex:  make(map[int64]*models.VaultNode),
		tagIndex:       make(map[string][]*models.VaultNode),
		regexes:        regexes,
	}
}

func (vc *VaultCrawler) Crawl() (*models.VaultNode, error) {
	root, err := vc.crawlDirectory(vc.RootPath, nil)
	if err != nil {
		return nil, err
	}

	vc.resolveLinks()

	return root, nil
}

func (vc *VaultCrawler) crawlDirectory(dirPath string, parent *models.VaultNode) (*models.VaultNode, error) {
	info, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}

	ID := generateIDFromPath(dirPath)
	URL := strings.TrimPrefix(dirPath, vc.RootPath)
	URL = utils.Slugify(URL)

	node := &models.VaultNode{
		ID:      ID,
		Path:    dirPath,
		URL:     URL,
		Name:    info.Name(),
		IsDir:   true,
		Size:    info.Size(),
		ModTime: info.ModTime(),
		Parent:  parent,
	}

	vc.fileIndex[dirPath] = node
	vc.idToNodeIndex[ID] = node

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(dirPath, entry.Name())

		if strings.HasPrefix(entry.Name(), ".") ||
			strings.HasPrefix(entry.Name(), "_") {
			continue
		}

		var childNode *models.VaultNode
		if entry.IsDir() {
			childNode, err = vc.crawlDirectory(entryPath, node)
		} else {
			childNode, err = vc.crawlFile(entryPath, node)
		}

		if err != nil {
			continue
		}

		node.Children = append(node.Children, childNode)
	}

	return node, nil
}

func (vc *VaultCrawler) crawlFile(filePath string, parent *models.VaultNode) (*models.VaultNode, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}

	ID := generateIDFromPath(filePath)

	ext := filepath.Ext(info.Name())
	baseName := strings.TrimSuffix(info.Name(), ext)
	URL := strings.TrimPrefix(filePath, vc.RootPath)

	if ext == ".md" {
		URL = strings.TrimSuffix(URL, filepath.Ext(URL)) + ".html"
		URL = utils.Slugify(URL)
		URL = path.Join("/", URL)
	}

	node := &models.VaultNode{
		ID:        ID,
		Path:      filePath,
		URL:       URL,
		Name:      info.Name(),
		Title:     baseName, // This may change during md processing if frontmatter has title
		BaseName:  baseName,
		Extension: ext,
		MimeType:  mime.TypeByExtension(ext),
		IsDir:     false,
		Size:      info.Size(),
		ModTime:   info.ModTime(),
		Parent:    parent,
	}

	var noteType models.NoteType

	if ext == ".md" {
		noteType = models.NoteTypeMarkdown
	} else if node.IsDir {
		noteType = models.NoteTypeFolder
	} else {
		noteType = models.NoteTypeUnknown
	}

	node.NoteType = noteType

	if hash, err := vc.calculateHash(filePath); err == nil {
		node.Hash = hash
	}

	vc.setIndexMaps(node)

	if noteType == models.NoteTypeMarkdown {
		vc.processMarkdownFile(node)
	}

	for _, tag := range node.Tags {
		if _, exists := vc.tagIndex[tag]; !exists {
			vc.tagIndex[tag] = make([]*models.VaultNode, 0)
			vc.tagIndex[tag] = append(vc.tagIndex[tag], node)
		} else {
			vc.tagIndex[tag] = append(vc.tagIndex[tag], node)
		}
	}
	return node, nil
}

func (vc *VaultCrawler) setIndexMaps(node *models.VaultNode) {
	vc.fileIndex[node.Path] = node
	vc.nameIndex[node.BaseName] = append(vc.nameIndex[node.BaseName], node)
	vc.idToNodeIndex[node.ID] = node

	cleanedDir := strings.TrimSuffix(strings.TrimPrefix(node.Path, vc.RootPath), node.Extension)
	subdirs := strings.Split(cleanedDir, string(filepath.Separator))
	for i := len(subdirs) - 1; i >= 0; i-- {
		dirSlice := subdirs[i:]
		key := strings.Join(dirSlice, string(filepath.Separator))
		key = vc.normalizeDir(key)

		vc.depthFileIndex[key] = node
	}

}

func (vc *VaultCrawler) GetTagIndex() map[string][]*models.VaultNode {
	return vc.tagIndex
}

func (vc *VaultCrawler) GetFileIndex() map[string]*models.VaultNode {
	return vc.fileIndex
}

func (vc *VaultCrawler) GetNameIndex() map[string][]*models.VaultNode {
	return vc.nameIndex
}

func (vc *VaultCrawler) GetFileDepthIndex() map[string]*models.VaultNode {
	return vc.depthFileIndex
}

func (vc *VaultCrawler) GetIdToNodeIndex() map[int64]*models.VaultNode {
	return vc.idToNodeIndex
}
