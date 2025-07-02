package generators

import (
	"encoding/json"
	"gobsidian/internal/config"
	"gobsidian/internal/models"
	"log"
	"strings"
)

type GraphGenerator struct {
	Config *config.Config
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

func (g *GraphGenerator) Generate(posts []*models.BlogPost, relativePathToPost map[string]*models.BlogPost, idToPost map[int64]*models.BlogPost) []byte {
	nodes := make([]*Node, 0)
	edges := make([]*Edge, 0)

	for _, post := range posts {
		nodes = append(nodes, &Node{ID: post.ID, Label: post.Title, URL: post.URL})

		for _, linkedPost := range post.LinkedTitlesStrings {
			if _, ok := relativePathToPost[linkedPost]; ok {
				edges = append(edges, &Edge{From: (*relativePathToPost[linkedPost]).ID, To: post.ID})
			} else {
				postID := g.handleLinksInsideFolders(linkedPost, relativePathToPost)
				if postID != 0 {
					edges = append(edges, &Edge{From: postID, To: post.ID})
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

func (g *GraphGenerator) handleLinksInsideFolders(relativePathText string, relativePathToPost map[string]*models.BlogPost) int64 {
	var lastPossibleMatch string
	for relativePath := range relativePathToPost {
		splitPath := strings.Split(relativePath, "/")
		if len(splitPath) > 0 {
			lastPossibleMatch = splitPath[len(splitPath)-1]

			if lastPossibleMatch == relativePathText {
				postFromMap, ok := relativePathToPost[relativePath]
				if ok {
					return postFromMap.ID
				}
			}

		}
	}
	return 0
}
