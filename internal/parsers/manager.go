package parsers

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/crawler"
	"gobsidian/internal/models"
	"gobsidian/internal/utils"
	"strings"
	"time"
)

type Manager struct {
	Parsers    map[models.NoteType]Parser
	SiteConfig config.SiteConfig
}

func NewManager(siteConfig config.SiteConfig, parsers map[models.NoteType]Parser) *Manager {
	return &Manager{
		Parsers:    parsers,
		SiteConfig: siteConfig,
	}
}

func (p *Manager) ParseNotes(rootNode *crawler.VaultNode) time.Duration {
	start := time.Now()

	currentNode := rootNode

	for _, node := range currentNode.Children {
		p.parseNode(node)
	}

	return time.Since(start)
}

func (p *Manager) parseNode(rootNode *crawler.VaultNode) {

}

// func (p *Manager) ParseNotes(
// 	rootNode *crawler.VaultNode,
// 	shallowNotes []*models.ScannedNote,
// 	notesRepository *repository.NoteRepository,
// ) time.Duration {
// 	start := time.Now()
// 	numWorkers := min(runtime.NumCPU(), len(shallowNotes))
// 	notesChan := make(chan *models.ScannedNote, len(shallowNotes))
// 	resultsChan := make(chan ParseResult, len(shallowNotes))

// 	var wg sync.WaitGroup
// 	for range numWorkers {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			for note := range notesChan {
// 				result := p.parseNoteWorker(note)
// 				resultsChan <- result
// 			}
// 		}()
// 	}

// 	for _, note := range shallowNotes {
// 		notesChan <- note
// 	}
// 	close(notesChan)

// 	wg.Wait()
// 	close(resultsChan)

// 	for result := range resultsChan {
// 		if result.Err != nil {
// 			continue
// 		}

// 		notesRepository.AddNote(result.Note)
// 	}

// 	return time.Since(start)
// }

func (p *Manager) parseNoteWorker(
	scannedNote *models.ScannedNote,
) ParseResult {
	if scannedNote == nil {
		return ParseResult{
			Note: nil,
			Err:  fmt.Errorf("received nil scannedNote"),
		}
	}

	noteType := scannedNote.GetNoteType()
	parsedNote, err := p.Parsers[noteType].Parse(scannedNote.FullPath, scannedNote.FileName)
	parsedNote.ScannedNote = scannedNote

	if err != nil {
		return ParseResult{
			Note: &parsedNote,
			Err:  err,
		}
	}

	if scannedNote.IsLandingPage {
		parsedNote.URL = "/"
	} else {
		URL := strings.TrimPrefix(scannedNote.FullPath, p.SiteConfig.InputDirectory)
		URL = strings.TrimPrefix(URL, "/")
		URL = strings.TrimSuffix(URL, ".md")

		slugifiedURL := "/" + utils.Slugify(URL) + ".html"
		parsedNote.URL = slugifiedURL
	}

	return ParseResult{
		Note: &parsedNote,
		Err:  nil,
	}
}
