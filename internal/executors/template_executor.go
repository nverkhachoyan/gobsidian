package executors

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"slices"
	"sort"
	"strings"
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
}

func NewTemplateExecutor(
	siteConfig *config.SiteConfig,
	logger *log.Logger,
	templates *template.Template,
) *TemplateExecutor {
	return &TemplateExecutor{
		SiteConfig: siteConfig,
		Logger:     logger,
		Templates:  templates,
	}
}

func (te *TemplateExecutor) Execute(
	notes []*models.ParsedNote,
	fileTree *models.Folder,
) (time.Duration, int, error) {
	start := time.Now()
	batchWriter := NewBatchWriter(te.SiteConfig.OutputDirectory, te.Logger, te.Templates)

	tags, postsByTag := te.computeTags(notes)
	tagsJSON, err := json.Marshal(tags)
	if err != nil {
		return time.Since(start), 0, fmt.Errorf("failed to marshal tags to JSON: %w", err)
	}
	baseData := te.createBaseTemplateData(fileTree, tags)

	for _, note := range notes {
		data := struct {
			Note         models.ParsedNote
			SiteTitle    string
			SiteSubtitle string
			BaseURL      string
			CurrentYear  int
			FileTree     *models.Folder
			Graph        template.JS
			AllTags      template.JS
			FileTreeJSON template.JS
		}{
			Note:         *note,
			SiteTitle:    baseData.SiteTitle,
			SiteSubtitle: baseData.SiteSubtitle,
			BaseURL:      baseData.BaseURL,
			CurrentYear:  baseData.CurrentYear,
			FileTree:     baseData.FileTree,
			Graph:        baseData.Graph,
			AllTags:      template.JS(tagsJSON),
			FileTreeJSON: baseData.FileTreeJSON,
		}

		filePath := note.URL
		if !strings.HasSuffix(filePath, ".html") {
			filePath += ".html"
		}
		cleanFilePath := strings.TrimPrefix(filePath, "/")

		batchWriter.AddFile(cleanFilePath, "post.html", data)

		previewPath := filepath.Join("previews", cleanFilePath)
		batchWriter.AddFile(previewPath, "preview.html", data)
	}

	for _, tag := range tags {
		if err := te.paginateTagPages(batchWriter, tag, postsByTag[tag.Slug], baseData); err != nil {
			return time.Since(start), 0, err
		}
	}

	if err := te.paginateIndexPages(batchWriter, notes, baseData); err != nil {
		return time.Since(start), 0, err
	}

	if err := te.addFolderPagesToBatchRecursive(batchWriter, fileTree, baseData); err != nil {
		return time.Since(start), 0, err
	}

	batchWriter.AddFile("404.html", "404.html", baseData)

	te.Logger.Debug("Writing all template files")
	if err := batchWriter.WriteAllFiles(); err != nil {
		return time.Since(start), 0, fmt.Errorf("failed to write template files: %w", err)
	}

	return time.Since(start), len(tags), nil
}

func (te *TemplateExecutor) addFolderPagesToBatchRecursive(
	batchWriter *BatchWriter,
	folder *models.Folder,
	baseData BaseTemplateData,
) error {
	if folder.Path != "" {
		if err := te.paginateFolderPages(batchWriter, folder, baseData); err != nil {
			return err
		}
	}

	for _, child := range folder.Children {
		if err := te.addFolderPagesToBatchRecursive(batchWriter, child, baseData); err != nil {
			return err
		}
	}

	return nil
}

func (te *TemplateExecutor) paginateIndexPages(
	batchWriter *BatchWriter,
	notes []*models.ParsedNote,
	baseData BaseTemplateData,
) error {
	if len(notes) == 0 {
		data := struct {
			BaseTemplateData
			Notes      []*models.ParsedNote
			Pagination PaginationData
			AllTags    []models.Tag
		}{
			BaseTemplateData: baseData,
			Notes:            notes,
			Pagination:       PaginationData{},
		}
		return batchWriter.AddFile("index.html", "index.html", data)
	}

	var landingPage *models.ParsedNote
	landingPageIndex := slices.IndexFunc(notes, func(note *models.ParsedNote) bool {
		return note.IsLandingPage
	})
	if landingPageIndex != -1 {
		landingPage = notes[landingPageIndex]
	}
	totalPages := (len(notes) + te.SiteConfig.NotesPerPage - 1) / te.SiteConfig.NotesPerPage

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * te.SiteConfig.NotesPerPage
		end := min(start+te.SiteConfig.NotesPerPage, len(notes))
		pageNotes := notes[start:end]

		var pagePath string
		if page == 1 {
			pagePath = "index.html"
		} else {
			pagePath = filepath.Join("page", fmt.Sprintf("%d", page), "index.html")
		}

		pagination := PaginationData{
			CurrentPage: page,
			TotalPages:  totalPages,
		}
		if page > 1 {
			pagination.HasPrev = true
			if page == 2 {
				pagination.PrevPageURL = te.SiteConfig.BaseURL
			} else {
				pagination.PrevPageURL = te.SiteConfig.BaseURL + fmt.Sprintf("page/%d/", page-1)
			}
		}
		if page < totalPages {
			pagination.HasNext = true
			pagination.NextPageURL = te.SiteConfig.BaseURL + fmt.Sprintf("page/%d/", page+1)
		}

		data := struct {
			BaseTemplateData
			Notes           []*models.ParsedNote
			LandingPageNote *models.ParsedNote
			Pagination      PaginationData
		}{
			BaseTemplateData: baseData,
			Notes:            pageNotes,
			LandingPageNote:  landingPage,
			Pagination:       pagination,
		}

		if err := batchWriter.AddFile(pagePath, "index.html", data); err != nil {
			return fmt.Errorf("failed to add index page %d: %w", page, err)
		}
	}

	return nil
}

func (te *TemplateExecutor) paginateTagPages(
	batchWriter *BatchWriter,
	tag models.Tag,
	posts []*models.ParsedNote,
	baseData BaseTemplateData,
) error {
	if len(posts) == 0 {
		return nil
	}

	totalPages := (len(posts) + te.SiteConfig.NotesPerPage - 1) / te.SiteConfig.NotesPerPage
	pathPrefix := filepath.Join("tag", tag.Slug)

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * te.SiteConfig.NotesPerPage
		end := min(start+te.SiteConfig.NotesPerPage, len(posts))
		pagePosts := posts[start:end]

		var pagePath string
		if page == 1 {
			pagePath = filepath.Join(pathPrefix, "index.html")
		} else {
			pagePath = filepath.Join(pathPrefix, "page", fmt.Sprintf("%d", page), "index.html")
		}

		pagination := PaginationData{
			CurrentPage: page,
			TotalPages:  totalPages,
		}
		if page > 1 {
			pagination.HasPrev = true
			if page == 2 {
				pagination.PrevPageURL = te.SiteConfig.BaseURL + strings.ReplaceAll(pathPrefix, "\\", "/") + "/"
			} else {
				pagination.PrevPageURL = te.SiteConfig.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(pathPrefix, "\\", "/"), page-1)
			}
		}
		if page < totalPages {
			pagination.HasNext = true
			pagination.NextPageURL = te.SiteConfig.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(pathPrefix, "\\", "/"), page+1)
		}

		data := struct {
			BaseTemplateData
			models.Tag
			Posts      []*models.ParsedNote
			Pagination PaginationData
		}{
			BaseTemplateData: baseData,
			Tag:              tag,
			Posts:            pagePosts,
			Pagination:       pagination,
		}

		if err := batchWriter.AddFile(pagePath, "tag.html", data); err != nil {
			return fmt.Errorf("failed to add tag page %s page %d: %w", tag.Slug, page, err)
		}
	}

	return nil
}

func (te *TemplateExecutor) paginateFolderPages(
	batchWriter *BatchWriter,
	folder *models.Folder,
	baseData BaseTemplateData,
) error {
	if len(folder.Posts) == 0 {
		data := struct {
			BaseTemplateData
			Folder     *models.Folder
			Posts      []*models.ParsedNote
			Pagination PaginationData
			AllTags    []models.Tag
		}{
			BaseTemplateData: baseData,
			Folder:           folder,
			Posts:            folder.Posts,
			Pagination:       PaginationData{},
		}
		folderPath := filepath.Join(folder.Path, "index.html")
		return batchWriter.AddFile(folderPath, "folder.html", data)
	}

	totalPages := (len(folder.Posts) + te.SiteConfig.NotesPerPage - 1) / te.SiteConfig.NotesPerPage

	for page := 1; page <= totalPages; page++ {
		start := (page - 1) * te.SiteConfig.NotesPerPage
		end := min(start+te.SiteConfig.NotesPerPage, len(folder.Posts))
		pagePosts := folder.Posts[start:end]

		var pagePath string
		if page == 1 {
			pagePath = filepath.Join(folder.Path, "index.html")
		} else {
			pagePath = filepath.Join(folder.Path, "page", fmt.Sprintf("%d", page), "index.html")
		}

		pagination := PaginationData{
			CurrentPage: page,
			TotalPages:  totalPages,
		}
		if page > 1 {
			pagination.HasPrev = true
			if page == 2 {
				pagination.PrevPageURL = te.SiteConfig.BaseURL + strings.ReplaceAll(folder.Path, "\\", "/") + "/"
			} else {
				pagination.PrevPageURL = te.SiteConfig.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(folder.Path, "\\", "/"), page-1)
			}
		}
		if page < totalPages {
			pagination.HasNext = true
			pagination.NextPageURL = te.SiteConfig.BaseURL + fmt.Sprintf("%s/page/%d/", strings.ReplaceAll(folder.Path, "\\", "/"), page+1)
		}

		data := struct {
			BaseTemplateData
			Folder     *models.Folder
			Posts      []*models.ParsedNote
			Pagination PaginationData
		}{
			BaseTemplateData: baseData,
			Folder:           folder,
			Posts:            pagePosts,
			Pagination:       pagination,
		}

		if err := batchWriter.AddFile(pagePath, "folder.html", data); err != nil {
			return fmt.Errorf("failed to add folder page %s page %d: %w", folder.Name, page, err)
		}
	}

	return nil
}

func (te *TemplateExecutor) createBaseTemplateData(
	fileTree *models.Folder,
	tags []models.Tag,
) BaseTemplateData {
	return BaseTemplateData{
		SiteTitle:    te.SiteConfig.SiteTitle,
		SiteSubtitle: te.SiteConfig.SiteSubtitle,
		BaseURL:      te.SiteConfig.BaseURL,
		CurrentYear:  time.Now().Year(),
		FileTree:     fileTree,
		Tags:         tags,
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
