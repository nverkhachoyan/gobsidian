package generators

import (
	"gobsidian/internal/models"
	"gobsidian/internal/utils"
	"strings"
)

type WikilinkResolver struct {
	notesByPath        map[string]*models.ParsedNote
	notesByFilename    map[string]*models.ParsedNote
	duplicateFilenames map[string][]*models.ParsedNote
	idToNote           map[int64]*models.ParsedNote
}

func NewWikilinkResolver() *WikilinkResolver {
	return &WikilinkResolver{
		notesByPath:        make(map[string]*models.ParsedNote),
		notesByFilename:    make(map[string]*models.ParsedNote),
		duplicateFilenames: make(map[string][]*models.ParsedNote),
		idToNote:           make(map[int64]*models.ParsedNote),
	}
}

func (wr *WikilinkResolver) AddNote(note *models.ParsedNote) {
	wr.notesByPath[note.RelativePath] = note
	wr.idToNote[note.ID] = note

	filename := note.FileName

	if existing, exists := wr.notesByFilename[filename]; exists {
		if existing != nil {
			wr.duplicateFilenames[filename] = []*models.ParsedNote{existing, note}
			delete(wr.notesByFilename, filename)
		} else {
			wr.duplicateFilenames[filename] = append(wr.duplicateFilenames[filename], note)
		}
	} else {
		wr.notesByFilename[filename] = note
	}
}

func (wr *WikilinkResolver) ResolveWikilink(linkText string) (*models.ParsedNote, bool) {
	if note, ok := wr.ResolveByPath(linkText + ".md"); ok {
		return note, true
	}

	if note, ok := wr.ResolveByFilename(linkText + ".md"); ok {
		return note, true
	}

	if note, ok := wr.ResolveByFilename(linkText); ok {
		return note, true
	}

	return nil, false
}

func (wr *WikilinkResolver) AddBacklink(targetNote *models.ParsedNote, sourceNote *models.ParsedNote, pseudoName string) {
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

func (wr *WikilinkResolver) ResolveByPath(path string) (*models.ParsedNote, bool) {
	note, ok := wr.notesByPath[path]
	return note, ok
}

func (wr *WikilinkResolver) ResolveByFilename(filename string) (*models.ParsedNote, bool) {
	note, ok := wr.notesByFilename[filename]
	return note, ok
}

func (wr *WikilinkResolver) GetAllByPath() map[string]*models.ParsedNote {
	return wr.notesByPath
}

func (wr *WikilinkResolver) GetAllByID() map[int64]*models.ParsedNote {
	return wr.idToNote
}
