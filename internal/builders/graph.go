package builders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gobsidian/internal/models"
	"gobsidian/internal/repository"

	"strings"

	"github.com/charmbracelet/log"
)

type GraphBuilder struct {
	Logger             *log.Logger
	outputDirectory    string
	generatedDirectory string
	fileName           string
	graph              *Graph
	notesRepository    *repository.NoteRepository
}

type Node struct {
	ID    int64  `json:"id"`
	Label string `json:"label"`
	URL   string `json:"url"`
}

type Edge struct {
	From int64 `json:"source"`
	To   int64 `json:"target"`
}

type Graph struct {
	Nodes []*Node `json:"nodes"`
	Edges []*Edge `json:"edges"`
}

func NewGraphBuilder(logger *log.Logger, outputDirectory string, generatedDirectory string, fileName string, notesRepository *repository.NoteRepository) *GraphBuilder {
	return &GraphBuilder{
		Logger:             logger,
		outputDirectory:    outputDirectory,
		generatedDirectory: generatedDirectory,
		fileName:           fileName,
		graph:              &Graph{},
		notesRepository:    notesRepository,
	}
}

func (g *GraphBuilder) Build(notes map[string]*models.ParsedNote) time.Duration {
	start := time.Now()
	notesByPath := g.notesRepository.GetAllByPath()

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

	g.graph.Nodes = nodes
	g.graph.Edges = edges

	endTime := time.Since(start)

	return endTime
}

func (g *GraphBuilder) handleLinksInsideFolders(relativePathText string, notesByPath map[string]*models.ParsedNote) int64 {
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

func (g *GraphBuilder) ExportAndSaveAsJSON() error {
	json, err := json.Marshal(g.graph)

	destPath := filepath.Join(g.outputDirectory, g.generatedDirectory, g.fileName)
	if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}
	if err := os.WriteFile(destPath, json, 0644); err != nil {
		return fmt.Errorf("%s: %w", g.fileName, err)
	}

	if err != nil {
		return err
	}
	return nil
}
