package repository

import (
	"gobsidian/internal/models"
	"gobsidian/internal/utils"
	"maps"
	"strings"
	"sync"
)

type NoteRepository struct {
	NotesByPath        map[string]*models.ParsedNote
	NotesByFilename    map[string]*models.ParsedNote
	DuplicateFilenames map[string][]*models.ParsedNote
	IdToNote           map[int64]*models.ParsedNote
	mu                 sync.RWMutex
}

func NewNoteRepository() *NoteRepository {
	return &NoteRepository{
		NotesByPath:        make(map[string]*models.ParsedNote),
		NotesByFilename:    make(map[string]*models.ParsedNote),
		DuplicateFilenames: make(map[string][]*models.ParsedNote),
		IdToNote:           make(map[int64]*models.ParsedNote),
	}
}

func (nr *NoteRepository) AddNote(note *models.ParsedNote) {
	nr.mu.Lock()
	defer nr.mu.Unlock()

	nr.NotesByPath[note.RelativePath] = note
	nr.IdToNote[note.ID] = note

	filename := note.FileName

	if existing, exists := nr.NotesByFilename[filename]; exists {
		if existing != nil {
			nr.DuplicateFilenames[filename] = []*models.ParsedNote{existing, note}
			delete(nr.NotesByFilename, filename)
		} else {
			nr.DuplicateFilenames[filename] = append(nr.DuplicateFilenames[filename], note)
		}
	} else {
		nr.NotesByFilename[filename] = note
	}
}

func (nr *NoteRepository) GetAllNotesSlice() []*models.ParsedNote {
	nr.mu.Lock()
	defer nr.mu.Unlock()

	notes := make([]*models.ParsedNote, 0, len(nr.NotesByPath))
	for _, note := range nr.NotesByPath {
		notes = append(notes, note)
	}
	return notes
}

func (wr *NoteRepository) AddBacklink(targetNote *models.ParsedNote, sourceNote *models.ParsedNote, pseudoName string) {
	wr.mu.Lock()
	defer wr.mu.Unlock()

	for _, link := range targetNote.LinkedFrom {
		if link.URL == sourceNote.URL {
			return // Already exists
		}
	}

	htmlFileName := utils.Slugify(strings.TrimSuffix(sourceNote.FileName, ".md")) + ".html"
	targetNote.LinkedFrom = append(targetNote.LinkedFrom, models.Link{
		Title:       sourceNote.Title,
		URL:         sourceNote.URL,
		RawFileName: htmlFileName,
		PseudoName:  pseudoName,
	})
}

func (tc *NoteRepository) ResolveWikilink(linkText string) (*models.ParsedNote, bool) {
	if note, ok := tc.GetByPath(linkText + ".md"); ok {
		return note, true
	}

	if note, ok := tc.GetByFilename(linkText + ".md"); ok {
		return note, true
	}

	if note, ok := tc.GetByFilename(linkText); ok {
		return note, true
	}

	return nil, false
}

func (tc *NoteRepository) GetAllByPath() map[string]*models.ParsedNote {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	result := make(map[string]*models.ParsedNote, len(tc.NotesByPath))
	maps.Copy(result, tc.NotesByPath)
	return result
}

func (tc *NoteRepository) GetAllByID() map[int64]*models.ParsedNote {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	result := make(map[int64]*models.ParsedNote, len(tc.IdToNote))
	maps.Copy(result, tc.IdToNote)
	return result
}

func (tc *NoteRepository) Count() int {
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	return len(tc.NotesByPath)
}

func (tc *NoteRepository) GetByPath(path string) (*models.ParsedNote, bool) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	note, ok := tc.NotesByPath[path]
	return note, ok
}

func (tc *NoteRepository) GetByFilename(filename string) (*models.ParsedNote, bool) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	note, ok := tc.NotesByFilename[filename]
	return note, ok
}

func (tc *NoteRepository) GetByID(id int64) (*models.ParsedNote, bool) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	note, ok := tc.IdToNote[id]
	return note, ok
}
