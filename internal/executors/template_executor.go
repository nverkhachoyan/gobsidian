package executors

import (
	"fmt"
	"path"
	"path/filepath"
	"time"

	"html/template"

	"gobsidian/internal/config"
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

type TemplateExecutor struct {
	SiteConfig *config.SiteConfig
	Logger     *log.Logger
	Templates  *template.Template
	rootNode   *models.VaultNode
	fileIndex  map[string]*models.VaultNode
	tagIndex   map[string][]*models.VaultNode
}

func NewTemplateExecutor(
	rootNode *models.VaultNode,
	fileIndex map[string]*models.VaultNode,
	tagIndex map[string][]*models.VaultNode,
	siteConfig *config.SiteConfig,
	logger *log.Logger,
	templates *template.Template,
) *TemplateExecutor {
	return &TemplateExecutor{
		rootNode:   rootNode,
		fileIndex:  fileIndex,
		tagIndex:   tagIndex,
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
	var landingPage *models.VaultNode
	var markdownNodesSlice []*models.VaultNode
	var folderNodesSlice []*models.VaultNode

	for _, node := range fileIndex {
		if node.IsLandingPageNote {
			landingPage = node
		}
		if node.GetNoteType() == models.NoteTypeMarkdown {
			markdownNodesSlice = append(markdownNodesSlice, node)
		} else if node.IsDir {
			folderNodesSlice = append(folderNodesSlice, node)
		}
	}

	te.executeHomePage(markdownNodesSlice, landingPage, batchWriter)
	te.executeMarkdowns(rootNode, batchWriter)
	te.executeFolders(folderNodesSlice, batchWriter)
	te.executeTagPages(te.tagIndex, batchWriter)

	te.Logger.Debug("Writing all template files")
	if err := batchWriter.WriteAllFiles(); err != nil {
		return time.Since(start), 0, fmt.Errorf("failed to write template files: %w", err)
	}

	return time.Since(start), 69, nil
}

func (te *TemplateExecutor) executeHomePage(
	markdownNodesSlice []*models.VaultNode,
	landingPageNode *models.VaultNode,
	batchWriter *BatchWriter) {
	renderableNodes := make([]*models.VaultNode, 0)
	for _, node := range markdownNodesSlice {
		if node.NoteType == models.NoteTypeMarkdown || node.NoteType == models.NoteTypeExcalidraw {
			renderableNodes = append(renderableNodes, node)
		}
	}

	numPages := 1
	if te.SiteConfig.NotesPerPage != 0 {
		numPages = len(renderableNodes) / te.SiteConfig.NotesPerPage
	}

	pages := make([][]*models.VaultNode, 0)
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
			LandingPageNote *models.VaultNode
			Notes           []*models.VaultNode
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
			data.LandingPageNote = nil
			pageDir := fmt.Sprintf("/pages/%d/index.html", i+1)
			batchWriter.AddFile(pageDir, "index.html", data)
		}

	}
}

func (te *TemplateExecutor) executeMarkdowns(rootNode *models.VaultNode, batchWriter *BatchWriter) {
	noteType := rootNode.GetNoteType()

	data := struct {
		Note         models.VaultNode
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

	switch noteType {
	case "markdown":
		batchWriter.AddFile(rootNode.URL, "post.html", data)
		previewUrl := filepath.Join("previews", rootNode.URL)
		batchWriter.AddFile(previewUrl, "preview.html", data)
	case "excalidraw":
		batchWriter.AddFile(rootNode.URL, "excalidraw.html", data)
		previewUrl := filepath.Join("previews", rootNode.URL)
		batchWriter.AddFile(previewUrl, "preview.html", data)
	default:
	}

	for _, child := range rootNode.Children {
		te.executeMarkdowns(child, batchWriter)
	}
}

func (te *TemplateExecutor) executeFolders(folderNodes []*models.VaultNode, batchWriter *BatchWriter) {
	for _, folderNode := range folderNodes {
		if folderNode.Path == te.SiteConfig.InputDirectory || !folderNode.IsDir {
			continue
		}

		renderableNodes := make([]*models.VaultNode, 0)
		subFolders := make([]*models.VaultNode, 0)
		for _, node := range folderNode.Children {
			if node.NoteType == models.NoteTypeMarkdown || node.NoteType == models.NoteTypeExcalidraw {
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

		pages := make([][]*models.VaultNode, 0)
		for i := 0; i < len(renderableNodes); i = i + te.SiteConfig.NotesPerPage {
			end := min(i+te.SiteConfig.NotesPerPage, len(renderableNodes))
			pages = append(pages, renderableNodes[i:end])
		}

		if len(pages) == 0 {
			data := struct {
				Folder       *models.VaultNode
				Subfolders   []*models.VaultNode
				Notes        []*models.VaultNode
				SiteTitle    string
				SiteSubtitle string
				BaseURL      string
				CurrentYear  int
				Pagination   PaginationData
			}{
				Folder:       folderNode,
				Subfolders:   subFolders,
				SiteTitle:    te.SiteConfig.SiteTitle,
				SiteSubtitle: te.SiteConfig.SiteSubtitle,
				BaseURL:      te.SiteConfig.BaseURL,
				CurrentYear:  time.Now().Year(),
				Pagination:   PaginationData{},
			}
			folderPath := path.Join(folderNode.URL, "index.html")
			batchWriter.AddFile(folderPath, "folder.html", data)
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
				Folder       *models.VaultNode
				Subfolders   []*models.VaultNode
				Notes        []*models.VaultNode
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

func (te *TemplateExecutor) executeTagPages(tagIndex map[string][]*models.VaultNode, batchWriter *BatchWriter) {
	for tagName, nodes := range tagIndex {
		renderableNodes := make([]*models.VaultNode, 0)
		for _, node := range nodes {
			if node.NoteType == models.NoteTypeMarkdown || node.NoteType == models.NoteTypeExcalidraw {
				renderableNodes = append(renderableNodes, node)
			}
		}

		numPages := 1
		if te.SiteConfig.NotesPerPage != 0 {
			numPages = len(renderableNodes) / te.SiteConfig.NotesPerPage
		}

		pages := make([][]*models.VaultNode, 0)
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
				pagination.NextPageURL = filepath.Join("/", "tag", tagName, "pages", fmt.Sprint(i+2))
			}

			if pagination.HasPrev && i == 1 {
				// note: for root node, we just give it base url
				pagination.PrevPageURL = te.SiteConfig.BaseURL
			} else if pagination.HasPrev {
				pagination.PrevPageURL = filepath.Join("/", "tag", tagName, "pages", fmt.Sprint(i))

			}

			data := struct {
				Tag          models.Tag
				Notes        []*models.VaultNode
				SiteTitle    string
				SiteSubtitle string
				BaseURL      string
				CurrentYear  int
				Pagination   PaginationData
			}{
				Tag: models.Tag{
					Name:  tagName,
					Slug:  tagName,
					Count: len(pageNodes),
				},
				Notes:        pageNodes,
				SiteTitle:    te.SiteConfig.SiteTitle,
				SiteSubtitle: te.SiteConfig.SiteSubtitle,
				BaseURL:      te.SiteConfig.BaseURL,
				CurrentYear:  time.Now().Year(),
				Pagination:   *pagination,
			}

			if i == 0 {
				tagPath := path.Join("/", "tag", tagName, "index.html")
				batchWriter.AddFile(tagPath, "tag.html", data)
			} else {
				tagPageDir := fmt.Sprintf("/tag/%s/pages/%d/index.html", tagName, i+1)
				batchWriter.AddFile(tagPageDir, "tag.html", data)
			}
		}

	}
}
