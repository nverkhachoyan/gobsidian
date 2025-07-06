package generators

import (
	"fmt"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

type NoteScanner struct {
	logger         *log.Logger
	inputDirectory string
	notesByPath    map[string]*models.ScannedNote
	uniqueID       int64
	mu             sync.RWMutex
}

func NewNoteScanner(logger *log.Logger, inputDirectory string) *NoteScanner {
	return &NoteScanner{
		logger:         logger,
		inputDirectory: inputDirectory,
		notesByPath:    make(map[string]*models.ScannedNote),
	}
}

func (ns *NoteScanner) AddNote(note *models.ScannedNote) {
	ns.mu.Lock()
	defer ns.mu.Unlock()

	ns.uniqueID = ns.uniqueID + 1
	note.ID = ns.uniqueID

	ns.notesByPath[note.RelativePath] = note

}

func (ns *NoteScanner) GetAllNotes() []*models.ScannedNote {
	ns.mu.RLock()
	defer ns.mu.RUnlock()

	notes := make([]*models.ScannedNote, 0, len(ns.notesByPath))
	for _, note := range ns.notesByPath {
		if note != nil {
			notes = append(notes, note)
		}
	}
	return notes
}

func (ns *NoteScanner) ScanAllNotes() (time.Duration, error) {
	start := time.Now()
	err := filepath.WalkDir(ns.inputDirectory, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			notePath := strings.TrimPrefix(path, ns.inputDirectory+"/")
			isInsideFolder := strings.Contains(notePath, "/")
			note := &models.ScannedNote{
				FileName:       info.Name(),
				FullPath:       path,
				RelativePath:   notePath,
				IsInsideFolder: isInsideFolder,
			}
			ns.AddNote(note)

		}
		return nil
	})

	endTime := time.Since(start)

	if err != nil {
		return endTime, fmt.Errorf("error during initial walk: %w", err)
	}

	return endTime, nil
}
