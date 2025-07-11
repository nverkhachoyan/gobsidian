package generators

import (
	"sort"

	"gobsidian/internal/models"
)

func (g *StaticSiteGenerator) sortNotes(notes []*models.ParsedNote) []*models.ParsedNote {
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].Date.After(notes[j].Date)
	})
	return notes
}
