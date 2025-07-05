package transformers

import (
	"gobsidian/internal/models"
	"gobsidian/internal/repository"
	"html/template"
	"sync"
)

type TransformContext struct {
	RenderedPostCache   map[string]template.HTML
	RenderingInProgress map[string]bool
	cacheMu             sync.RWMutex
	notesRepository     *repository.NoteRepository
	EmbeddedPosts       map[string]models.EmbeddedPost

	// Transform Config
	IsDarkMode      bool
	BaseURL         string
	OutputDirectory string
}

func NewTransformContext(baseURL, outputDir string, notesRepository *repository.NoteRepository) *TransformContext {
	return &TransformContext{
		RenderedPostCache:   make(map[string]template.HTML),
		RenderingInProgress: make(map[string]bool),
		notesRepository:     notesRepository,
		EmbeddedPosts:       make(map[string]models.EmbeddedPost),
		BaseURL:             baseURL,
		IsDarkMode:          true,
		OutputDirectory:     outputDir,
	}
}

func (tc *TransformContext) getCachedPost(url string) (template.HTML, bool) {
	tc.cacheMu.RLock()
	defer tc.cacheMu.RUnlock()
	post, ok := tc.RenderedPostCache[url]
	return post, ok
}

func (tc *TransformContext) setCachedPost(url string, content template.HTML) {
	tc.cacheMu.Lock()
	defer tc.cacheMu.Unlock()
	tc.RenderedPostCache[url] = content
}

func (tc *TransformContext) isRenderingInProgress(url string) bool {
	tc.cacheMu.RLock()
	defer tc.cacheMu.RUnlock()
	return tc.RenderingInProgress[url]
}

func (tc *TransformContext) setRenderingInProgress(url string, inProgress bool) {
	tc.cacheMu.Lock()
	defer tc.cacheMu.Unlock()
	tc.RenderingInProgress[url] = inProgress
}
