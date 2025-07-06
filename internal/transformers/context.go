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
	cacheMu             *sync.RWMutex
	notesRepository     *repository.NoteRepository
	EmbeddedPosts       map[string]models.EmbeddedPost

	IsDarkMode      bool
	BaseURL         string
	OutputDirectory string
}

func NewTransformContext(baseURL, outputDir string, notesRepository *repository.NoteRepository) *TransformContext {
	return &TransformContext{
		RenderedPostCache:   make(map[string]template.HTML),
		RenderingInProgress: make(map[string]bool),
		cacheMu:             &sync.RWMutex{},
		notesRepository:     notesRepository,
		EmbeddedPosts:       make(map[string]models.EmbeddedPost),
		BaseURL:             baseURL,
		IsDarkMode:          true,
		OutputDirectory:     outputDir,
	}
}

// To ensure each note context has its data structures and mutex
func (tc *TransformContext) Clone() *TransformContext {
	return &TransformContext{
		RenderedPostCache:   make(map[string]template.HTML),
		RenderingInProgress: make(map[string]bool),
		EmbeddedPosts:       make(map[string]models.EmbeddedPost),
		cacheMu:             &sync.RWMutex{},
		notesRepository:     tc.notesRepository,
		BaseURL:             tc.BaseURL,
		IsDarkMode:          tc.IsDarkMode,
		OutputDirectory:     tc.OutputDirectory,
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
	if !inProgress {
		delete(tc.RenderingInProgress, url)
	} else {
		tc.RenderingInProgress[url] = true
	}
}
