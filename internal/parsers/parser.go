package parsers

import (
	"gobsidian/internal/models"
)

type Parser interface {
	Parse(filePath string, fileName string) (models.ParsedNote, error)
}
