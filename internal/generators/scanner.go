package generators

import (
	"gobsidian/internal/models"
	"sync"
)

type NoteScanner struct {
	inputDirectory string
	uniqueID       int64

	NotesByPath map[string]*models.ShallowNote

	mu sync.RWMutex
}

func NewNoteScanner(inputDirectory string) *NoteScanner {
	return &NoteScanner{
		inputDirectory: inputDirectory,

		NotesByPath: make(map[string]*models.ShallowNote),
	}
}

func (ns *NoteScanner) AddNote(note *models.ShallowNote) {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	ns.uniqueID = ns.uniqueID + 1
	note.ID = ns.uniqueID

	ns.NotesByPath[note.RelativePath] = note

}

func (ns *NoteScanner) GetAllNotes() []*models.ShallowNote {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	notes := make([]*models.ShallowNote, 0, len(ns.NotesByPath))
	for _, note := range ns.NotesByPath {
		if note != nil {
			notes = append(notes, note)
		}
	}
	return notes
}
