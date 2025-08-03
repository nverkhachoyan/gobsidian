package executors

import (
	"fmt"
	"path"
	"path/filepath"
	"sort"
	"time"

	"html/template"

	"gobsidian/internal/config"
	"gobsidian/internal/crawler"
	"gobsidian/internal/models"

	"github.com/charmbracelet/log"
)

type PaginationData struct {
	HasPrev     bool
	PrevPageURL string
	HasNext     bool
	NextPageURL string
	CurrentPage int
	TotalPages  int
}

type BaseTemplateData struct {
	SiteTitle    string
	SiteSubtitle string
	BaseURL      string
	CurrentYear  int
	FileTree     *models.Folder
	Graph        template.JS
	Tags         []models.Tag
	AllTags      template.JS
	FileTreeJSON template.JS
}

type TemplateExecutor struct {
	SiteConfig *config.SiteConfig
	Logger     *log.Logger
	Templates  *template.Template
	rootNode   *crawler.VaultNode
	fileIndex  map[string]*crawler.VaultNode
}

func NewTemplateExecutor(
	rootNode *crawler.VaultNode,
	fileIndex map[string]*crawler.VaultNode,
	siteConfig *config.SiteConfig,
	logger *log.Logger,
	templates *template.Template,
) *TemplateExecutor {
	return &TemplateExecutor{
		rootNode:   rootNode,
		fileIndex:  fileIndex,
		SiteConfig: siteConfig,
		Logger:     logger,
		Templates:  templates,
	}
}

func (te *TemplateExecutor) Execute() (time.Duration, int, error) {
	start := time.Now()
	rootNode := te.rootNode
	fileIndex := te.fileIndex
	batchWriter := NewBatchWriter(te.SiteConfig.OutputDirectory, te.Logger, te.Templates)
	var landingPage *crawler.VaultNode
	var markdownNodesSlice []*crawler.VaultNode
	var folderNodesSlice []*crawler.VaultNode

	for _, node := range fileIndex {
		if node.IsLandingPageNote {
			landingPage = node
		}

		if node.GetNoteType() == crawler.NoteTypeMarkdown {
			markdownNodesSlice = append(markdownNodesSlice, node)
		} else if node.IsDir {
			folderNodesSlice = append(folderNodesSlice, node)
		}
	}

	te.executeHomePage(markdownNodesSlice, landingPage, batchWriter)
	te.executeMarkdowns(rootNode, batchWriter)
	te.executeFolders(folderNodesSlice, batchWriter)

	te.Logger.Debug("Writing all template files")
	if err := batchWriter.WriteAllFiles(); err != nil {
		return time.Since(start), 0, fmt.Errorf("failed to write template files: %w", err)
	}

	return time.Since(start), 69, nil
}

func (te *TemplateExecutor) executeHomePage(
	markdownNodesSlice []*crawler.VaultNode,
	landingPageNode *crawler.VaultNode,
	batchWriter *BatchWriter) {
	renderableNodes := make([]*crawler.VaultNode, 0)
	for _, node := range markdownNodesSlice {
		if node.NoteType == crawler.NoteTypeMarkdown || node.NoteType == crawler.NoteTypeExcalidraw {
			renderableNodes = append(renderableNodes, node)
		}
	}

	numPages := 1
	if te.SiteConfig.NotesPerPage != 0 {
		numPages = len(renderableNodes) / te.SiteConfig.NotesPerPage
	}

	pages := make([][]*crawler.VaultNode, 0)
	for i := 0; i < len(renderableNodes); i = i + te.SiteConfig.NotesPerPage {
		end := min(i+te.SiteConfig.NotesPerPage, len(renderableNodes))
		pages = append(pages, renderableNodes[i:end])
	}

	for i, pageNodes := range pages {
		pagination := &PaginationData{
			HasPrev:     i > 0,
			HasNext:     len(pages) > i+1,
			CurrentPage: i + 1,
			TotalPages:  numPages,
		}

		if pagination.HasNext {
			pagination.NextPageURL = filepath.Join("/", "pages", fmt.Sprint(i+2))
		}

		if pagination.HasPrev && i == 1 {
			// note: for root node, we just give it base url
			pagination.PrevPageURL = te.SiteConfig.BaseURL
		} else if pagination.HasPrev {
			pagination.PrevPageURL = filepath.Join("/", "pages", fmt.Sprint(i))

		}

		data := struct {
			LandingPageNote *crawler.VaultNode
			Notes           []*crawler.VaultNode
			SiteTitle       string
			SiteSubtitle    string
			BaseURL         string
			CurrentYear     int
			Pagination      PaginationData
		}{
			LandingPageNote: landingPageNode,
			Notes:           pageNodes,
			SiteTitle:       te.SiteConfig.SiteTitle,
			SiteSubtitle:    te.SiteConfig.SiteSubtitle,
			BaseURL:         te.SiteConfig.BaseURL,
			CurrentYear:     time.Now().Year(),
			Pagination:      *pagination,
		}

		if i == 0 {
			batchWriter.AddFile("index.html", "index.html", data)
		} else {
			pageDir := fmt.Sprintf("/pages/%d/index.html", i+1)
			batchWriter.AddFile(pageDir, "index.html", data)
		}

	}
}

func (te *TemplateExecutor) executeMarkdowns(rootNode *crawler.VaultNode, batchWriter *BatchWriter) {
	if rootNode.GetNoteType() == crawler.NoteTypeMarkdown {
		data := struct {
			Note         crawler.VaultNode
			SiteTitle    string
			SiteSubtitle string
			BaseURL      string
			CurrentYear  int
		}{
			Note:         *rootNode,
			SiteTitle:    te.SiteConfig.SiteTitle,
			SiteSubtitle: te.SiteConfig.SiteSubtitle,
			BaseURL:      te.SiteConfig.BaseURL,
			CurrentYear:  time.Now().Year(),
		}

		batchWriter.AddFile(rootNode.URL, "post.html", data)
		previewUrl := filepath.Join("previews", rootNode.URL)
		batchWriter.AddFile(previewUrl, "preview.html", data)
	}

	for _, child := range rootNode.Children {
		te.executeMarkdowns(child, batchWriter)
	}
}

func (te *TemplateExecutor) executeFolders(folderNodes []*crawler.VaultNode, batchWriter *BatchWriter) {
	for _, folderNode := range folderNodes {
		if folderNode.Path == te.SiteConfig.InputDirectory || !folderNode.IsDir {
			continue
		}

		renderableNodes := make([]*crawler.VaultNode, 0)
		subFolders := make([]*crawler.VaultNode, 0)
		for _, node := range folderNode.Children {
			if node.NoteType == crawler.NoteTypeMarkdown || node.NoteType == crawler.NoteTypeExcalidraw {
				renderableNodes = append(renderableNodes, node)
			}
			if node.IsDir {
				subFolders = append(subFolders, node)
			}
		}

		numPages := 1
		if te.SiteConfig.NotesPerPage != 0 {
			numPages = len(renderableNodes) / te.SiteConfig.NotesPerPage
		}

		pages := make([][]*crawler.VaultNode, 0)
		for i := 0; i < len(renderableNodes); i = i + te.SiteConfig.NotesPerPage {
			end := min(i+te.SiteConfig.NotesPerPage, len(renderableNodes))
			pages = append(pages, renderableNodes[i:end])
		}

		for i, pageNodes := range pages {
			pagination := &PaginationData{
				HasPrev:     i > 0,
				HasNext:     len(pages) > i+1,
				CurrentPage: i + 1,
				TotalPages:  numPages,
			}

			if pagination.HasNext {
				pagination.NextPageURL = path.Join(te.SiteConfig.BaseURL, folderNode.URL, "pages", fmt.Sprint(i+2))
			}

			if pagination.HasPrev && i == 1 {
				pagination.PrevPageURL = filepath.Clean("/" + folderNode.URL)
			} else if pagination.HasPrev {
				pagination.PrevPageURL = filepath.Clean(filepath.Join(te.SiteConfig.BaseURL, folderNode.URL, "pages", fmt.Sprint(i)))
			}

			data := struct {
				Folder       *crawler.VaultNode
				Subfolders   []*crawler.VaultNode
				Notes        []*crawler.VaultNode
				SiteTitle    string
				SiteSubtitle string
				BaseURL      string
				CurrentYear  int
				Pagination   PaginationData
			}{
				Folder:       folderNode,
				Subfolders:   subFolders,
				Notes:        pageNodes,
				SiteTitle:    te.SiteConfig.SiteTitle,
				SiteSubtitle: te.SiteConfig.SiteSubtitle,
				BaseURL:      te.SiteConfig.BaseURL,
				CurrentYear:  time.Now().Year(),
				Pagination:   *pagination,
			}

			if i == 0 {
				folderPath := path.Join(folderNode.URL, "index.html")
				batchWriter.AddFile(folderPath, "folder.html", data)
			} else {
				folderPath := path.Join(folderNode.URL, "pages", fmt.Sprint(i+1), "index.html")
				batchWriter.AddFile(folderPath, "folder.html", data)
			}

		}

	}
}

func (te *TemplateExecutor) computeTags(
	notes []*models.ParsedNote,
) ([]models.Tag, map[string][]*models.ParsedNote) {
	tagsMap := make(map[string]*models.Tag)
	for _, note := range notes {
		for _, tag := range note.Tags {
			tagsMap[tag.Slug] = &tag
			tagsMap[tag.Slug].Count++
		}
	}

	tags := make([]models.Tag, 0, len(tagsMap))
	for _, tag := range tagsMap {
		tags = append(tags, *tag)
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Name < tags[j].Name
	})

	postsByTag := make(map[string][]*models.ParsedNote, len(tags))

	for _, note := range notes {
		for _, tag := range note.Tags {
			postsByTag[tag.Slug] = append(postsByTag[tag.Slug], note)
		}
	}

	return tags, postsByTag
}
