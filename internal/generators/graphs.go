package generators

import (
	"encoding/json"

	"gobsidian/internal/models"

	"strings"

	"github.com/charmbracelet/log"
)

type GraphGenerator struct {
	Logger *log.Logger
}

type Node struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Edge struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

type Graph struct {
	Nodes []*Node `json:"nodes"`
	Edges []*Edge `json:"edges"`
}

func NewGraphGenerator(cfg *GraphGenerator) *GraphGenerator {
	return &GraphGenerator{
		Logger: cfg.Logger,
	}
}

func (g *GraphGenerator) Generate(notes map[string]*models.ParsedNote, wikilinkResolver *WikilinkResolver) []byte {
	notesByPath := wikilinkResolver.GetAllByPath()

	nodes := make([]*Node, 0)
	edges := make([]*Edge, 0)

	for _, note := range notes {
		nodes = append(nodes, &Node{ID: note.ID, Label: note.Title, URL: note.URL})

		for _, linkedPost := range note.Wikilinks {
			if _, ok := notesByPath[linkedPost]; ok {
				edges = append(edges, &Edge{From: (*notesByPath[linkedPost]).ID, To: note.ID})
			} else {
				postID := g.handleLinksInsideFolders(linkedPost, notesByPath)
				if postID != 0 {
					edges = append(edges, &Edge{From: postID, To: note.ID})
				}
			}
		}

	}

	json, err := json.Marshal(Graph{
		Nodes: nodes,
		Edges: edges,
	})

	if err != nil {
		g.Logger.Print("Failed to generate JSON", "error", err)
		return nil
	}

	return json
}

func (g *GraphGenerator) handleLinksInsideFolders(relativePathText string, notesByPath map[string]*models.ParsedNote) int64 {
	var lastPossibleMatch string
	for relativePath := range notesByPath {
		splitPath := strings.Split(relativePath, "/")
		if len(splitPath) > 0 {
			lastPossibleMatch = splitPath[len(splitPath)-1]

			if lastPossibleMatch == relativePathText {
				postFromMap, ok := notesByPath[relativePath]
				if ok {
					return postFromMap.ID
				}
			}

		}
	}
	return 0
}
