package repository

import "gobsidian/internal/models"

type Repository interface {
	AddNote(note *models.ParsedNote)
	GetByPath(path string) (*models.ParsedNote, bool)
	GetByFilename(filename string) (*models.ParsedNote, bool)
	GetByID(id int64) (*models.ParsedNote, bool)
	GetAllByPath() map[string]*models.ParsedNote
	GetAllByID() map[int64]*models.ParsedNote
	ResolveWikilink(linkText string) (*models.ParsedNote, bool)
	Count() int
}
