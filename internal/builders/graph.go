package builders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gobsidian/internal/crawler"

	"github.com/charmbracelet/log"
)

type GraphBuilder struct {
	Logger             *log.Logger
	outputDirectory    string
	generatedDirectory string
	fileName           string
	graph              *Graph
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

func NewGraphBuilder(logger *log.Logger, outputDirectory string, generatedDirectory string, fileName string) *GraphBuilder {
	return &GraphBuilder{
		Logger:             logger,
		outputDirectory:    outputDirectory,
		generatedDirectory: generatedDirectory,
		fileName:           fileName,
		graph:              &Graph{},
	}
}

func (g *GraphBuilder) Build(fileIndex map[string]*crawler.VaultNode, IDToNodeIndex map[int64]*crawler.VaultNode) time.Duration {
	start := time.Now()

	nodes := make([]*Node, 0)
	edges := make([]*Edge, 0)

	for _, node := range fileIndex {
		if node.NoteType != crawler.NoteTypeMarkdown {
			continue
		}

		nodes = append(nodes, &Node{ID: node.ID, Label: node.Title, URL: node.URL})
		for _, backlink := range node.Backlinks {
			edges = append(edges, &Edge{
				From: backlink.TargetID,
				To:   node.ID,
			})

		}
	}

	g.graph.Nodes = nodes
	g.graph.Edges = edges

	return time.Since(start)
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
