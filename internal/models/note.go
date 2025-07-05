package models

import (
	"html/template"
	"time"
)

type Frontmatter struct {
	Title     string `yaml:"title"`
	Date      string `yaml:"date"`
	Author    string `yaml:"author"`
	UpdatedAt string `yaml:"updatedAt"`
}

type ScannedNote struct {
	ID             int64
	FileName       string
	RelativePath   string
	FullPath       string
	IsInsideFolder bool
}

type ParsedNote struct {
	*ScannedNote

	// Parsed content
	Title       string
	RawBody     []byte
	HTMLContent template.HTML
	Date        time.Time
	Author      string
	UpdatedAt   *time.Time
	Images      []Image
	Tags        []Tag
	LinkedFrom  []Link
	Wikilinks   []string
	URL         string
}

type EmbeddedPost struct {
	ID      string
	Content template.HTML
}

type Image struct {
	RelativePath string
	Alt          string
	Width        string
	Height       string
}

type Link struct {
	Title       string
	URL         string
	RawFileName string
	PseudoName  string
}

type Tag struct {
	Name string
	Slug string
}

type Folder struct {
	ID       int64
	Name     string
	Path     string
	Posts    []*ParsedNote
	Children map[string]*Folder
}
