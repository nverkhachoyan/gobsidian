package scanner

import (
	"fmt"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

type NoteScanner struct {
	logger      *log.Logger
	siteConfig  *config.SiteConfig
	notesByPath map[string]*models.ScannedNote
	uniqueID    int64
	mu          sync.RWMutex
}

func NewNoteScanner(logger *log.Logger, siteConfig *config.SiteConfig) *NoteScanner {
	return &NoteScanner{
		logger:      logger,
		siteConfig:  siteConfig,
		notesByPath: make(map[string]*models.ScannedNote),
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
	err := filepath.WalkDir(ns.siteConfig.InputDirectory, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			isLandingPage := false
			notePath := strings.TrimPrefix(path, ns.siteConfig.InputDirectory+"/")
			notePath = strings.TrimSuffix(notePath, ".md")
			isInsideFolder := strings.Contains(notePath, "/")
			sanitizedName := strings.TrimSuffix(info.Name(), ".md")

			if !isInsideFolder && sanitizedName == ns.siteConfig.LandingPage {
				isLandingPage = true
			}

			note := &models.ScannedNote{
				FileName:       sanitizedName,
				FullPath:       path,
				RelativePath:   notePath,
				IsInsideFolder: isInsideFolder,
				IsLandingPage:  isLandingPage,
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

func (ns *NoteScanner) ScanNotesByPaths(paths []string) (time.Duration, error) {
	start := time.Now()

	for _, path := range paths {

		noteFile, err := os.Open(path)
		if err != nil {
			return time.Duration(0), fmt.Errorf("error opening note file: %w", err)
		}
		defer noteFile.Close()

		notePath := strings.TrimPrefix(path, ns.siteConfig.InputDirectory+"/")
		isInsideFolder := strings.Contains(notePath, "/")

		_, ok := ns.notesByPath[notePath]
		if ok {
			delete(ns.notesByPath, notePath)
		}

		note := &models.ScannedNote{
			FileName:       filepath.Base(path),
			FullPath:       path,
			RelativePath:   notePath,
			IsInsideFolder: isInsideFolder,
		}

		ns.AddNote(note)
	}

	endTime := time.Since(start)

	return endTime, nil
}
