package builders

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gobsidian/internal/models"

	"github.com/charmbracelet/log"
)

type GraphType string

const (
	GraphTypeGlobal GraphType = "GlobalGraph"
	GraphTypeLocal  GraphType = "LocalGraph"
)

type GraphBuilder struct {
	Logger             *log.Logger
	outputDirectory    string
	generatedDirectory string
	fileName           string
	graph              *Graph
	localGraphs        map[string]*Graph
	landingPageGraph   *Graph
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
	IsForHomePage bool   `json:"-"`
	Nodes         []Node `json:"nodes"`
	Edges         []Edge `json:"edges"`
}

func NewGraphBuilder(logger *log.Logger, outputDirectory string, generatedDirectory string, fileName string) *GraphBuilder {
	return &GraphBuilder{
		Logger:             logger,
		outputDirectory:    outputDirectory,
		generatedDirectory: generatedDirectory,
		fileName:           fileName,
		graph:              &Graph{},
		localGraphs:        make(map[string]*Graph, 0),
	}
}

func (g *GraphBuilder) Build(fileIndex map[string]*models.VaultNode) time.Duration {
	start := time.Now()

	nodes := make([]Node, 0)
	edges := make([]Edge, 0)

	for _, node := range fileIndex {
		if node.NoteType != models.NoteTypeMarkdown && node.NoteType != models.NoteTypeExcalidraw {
			continue
		}

		nodes = append(nodes, Node{ID: node.ID, Label: node.Title, URL: node.URL})
		for _, backlink := range node.Backlinks {
			edges = append(edges, Edge{
				From: backlink.TargetID,
				To:   node.ID,
			})

		}

		g.buildLocalGraphs(node)
	}

	g.graph.Nodes = nodes
	g.graph.Edges = edges

	return time.Since(start)
}

func (g *GraphBuilder) buildLocalGraphs(rootNode *models.VaultNode) {
	if rootNode.NoteType != models.NoteTypeMarkdown && rootNode.NoteType != models.NoteTypeExcalidraw {
		return
	}

	isLandingPage := rootNode.IsLandingPageNote

	if _, exists := g.localGraphs[rootNode.Path]; !exists {
		g.localGraphs[rootNode.URL] = &Graph{
			IsForHomePage: isLandingPage,
			Nodes:         make([]Node, 0),
			Edges:         make([]Edge, 0),
		}
	}
	localGraph := g.localGraphs[rootNode.URL]
	addedNodes := make(map[int64]bool)

	localGraph.Nodes = append(localGraph.Nodes, Node{ID: rootNode.ID, Label: rootNode.Title, URL: rootNode.URL})
	addedNodes[rootNode.ID] = true

	for _, backlink := range rootNode.Backlinks {
		if _, exists := addedNodes[backlink.TargetID]; !exists {
			localGraph.Nodes = append(localGraph.Nodes, Node{ID: backlink.TargetID, Label: backlink.TargetNode.Title, URL: backlink.TargetNode.URL})
			addedNodes[backlink.TargetID] = true
		}

		localGraph.Edges = append(localGraph.Edges, Edge{
			From: backlink.TargetID,
			To:   rootNode.ID,
		})
	}

	for _, link := range rootNode.Links {
		if link.TargetNode == nil {
			continue
		}
		if _, exists := addedNodes[link.TargetID]; !exists {
			localGraph.Nodes = append(localGraph.Nodes, Node{ID: link.TargetID, Label: link.TargetNode.Title, URL: link.TargetNode.URL})
			addedNodes[link.TargetID] = true
		}

		localGraph.Edges = append(localGraph.Edges, Edge{
			From: rootNode.ID,
			To:   link.TargetID,
		})
	}
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

func (g *GraphBuilder) ExportAndSaveLocalGraphsAsJSON() error {
	for url, graph := range g.localGraphs {
		json, err := json.Marshal(graph)

		if err != nil {
			return err
		}

		urlWithoutExt := strings.TrimRight(url, filepath.Ext(url))
		fileName := fmt.Sprintf("%s.json", urlWithoutExt)

		if graph.IsForHomePage {
			pathWithoutName := filepath.Dir(fileName)
			fileName = filepath.Join(pathWithoutName, "gobsidian-homepagegraph.json")
		}

		destPath := filepath.Join(g.outputDirectory, g.generatedDirectory, "graphs", fileName)
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		if err := os.WriteFile(destPath, json, 0644); err != nil {
			return fmt.Errorf("%s: %w", g.fileName, err)
		}
	}
	return nil
}
