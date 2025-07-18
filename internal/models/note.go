package models

import (
	"fmt"
	"html/template"
	"strings"
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
	IsLandingPage  bool
}

type ParsedNote struct {
	*ScannedNote

	Title             string
	RawBody           []byte
	HTMLContent       template.HTML
	Date              time.Time
	Author            string
	UpdatedAt         *time.Time
	Images            []Image
	Tags              []Tag
	LinkedFrom        []Link
	Wikilinks         []string
	URL               string
	Warnings          []string
	MissingFiles      []string
	CssClasses        []string
	Footnotes         []Footnote
	EmbeddedFootnotes []Footnote
	Breadcrumbs       []Breadcrumb
}

type Breadcrumb struct {
	Title  string
	URL    string
	IsLast bool
}

type Footnote struct {
	ID     string
	Number int
	Text   string
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
	FullPath    string
	RawFileName string
	PseudoName  string
}

type Tag struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Count int    `json:"count"`
}

type Folder struct {
	ID       int64
	Name     string
	Path     string
	Posts    []*ParsedNote
	Children map[string]*Folder
}

func (p *ParsedNote) HasWarnings() bool {
	return len(p.Warnings) > 0
}

func (p *ParsedNote) HasMissingFiles() bool {
	return len(p.MissingFiles) > 0
}

func (p *ParsedNote) GetWarningsString() string {
	if !p.HasWarnings() {
		return ""
	}
	return fmt.Sprintf("warnings: %s", strings.Join(p.Warnings, "; "))
}

func (p *ParsedNote) GetMissingFilesString() string {
	if !p.HasMissingFiles() {
		return ""
	}
	return fmt.Sprintf("missing files: %s", strings.Join(p.MissingFiles, ", "))
}
