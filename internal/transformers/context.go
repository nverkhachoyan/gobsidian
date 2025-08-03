package transformers

import (
	"gobsidian/internal/models"
	"html/template"
	"sync"
)

type TransformContext struct {
	RenderedPostCache   map[string]template.HTML
	RenderingInProgress map[string]bool
	cacheMu             *sync.RWMutex
	EmbeddedPosts       map[string]models.EmbeddedPost

	BaseURL         string
	Port            int
	OutputDirectory string
}

func NewTransformContext(baseURL string, port int, outputDir string) *TransformContext {
	return &TransformContext{
		RenderedPostCache:   make(map[string]template.HTML),
		RenderingInProgress: make(map[string]bool),
		cacheMu:             &sync.RWMutex{},
		EmbeddedPosts:       make(map[string]models.EmbeddedPost),
		BaseURL:             baseURL,
		Port:                port,
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
		BaseURL:             tc.BaseURL,
		Port:                tc.Port,
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
