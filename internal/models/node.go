package models

import (
	"html/template"
	"slices"
	"strings"
	"time"
)

type NoteType string

const (
	NoteTypeMarkdown   NoteType = "markdown"
	NoteTypeExcalidraw NoteType = "excalidraw"
	NoteTypeFolder     NoteType = "folder"
	NoteTypeUnknown    NoteType = "unknown"
)

type GenericFrontmatter map[string]any

type VaultNode struct {
	ID        int64        `json:"id"`
	Path      string       `json:"path"`
	Name      string       `json:"name"`
	BaseName  string       `json:"baseName"` // name without extension
	Extension string       `json:"extension"`
	MimeType  string       `json:"mimeType"`
	IsDir     bool         `json:"isDir"`
	Size      int64        `json:"size"`
	ModTime   time.Time    `json:"modTime"`
	Hash      string       `json:"hash,omitempty"` // MD5 hash for files
	Children  []*VaultNode `json:"children,omitempty"`
	Parent    *VaultNode   `json:"-"`

	IsLandingPageNote  bool          `json:"isLandingPageNote"`
	NoteType           NoteType      `json:"noteType"`
	Markdown           string        `json:"markdown"`
	HTML               template.HTML `json:"HTML"`
	Tags               []string      `json:"tags,omitempty"`
	Links              []Link        `json:"links,omitempty"`
	Backlinks          []Link        `json:"backlinks,omitempty"`
	EmbeddedFootnotes  []Footnote
	WordCount          int                `json:"wordCount,omitempty"`
	Title              string             `json:"title"`
	URL                string             `json:"url"`
	Author             string             `json:"author"`
	Date               time.Time          `json:"date"`
	UpdatedAt          time.Time          `json:"updatedAt"`
	Breadcrumbs        []Breadcrumb       `json:"breadcrumbs"`
	Footnotes          []Footnote         `json:"footnotes"`
	GenericFrontmatter GenericFrontmatter `json:"frontmatter"`
	CssClasses         []string

	Excalidraw        *Excalidraw
	ExcalidrawPayload ExcalidrawPayload
}

type Breadcrumb struct {
	Title  string
	URL    string
	IsLast bool
}

func (vn *VaultNode) GetNoteType() NoteType {
	return vn.NoteType
}

func (vn *VaultNode) IsImageFile() bool {
	imageMimeTypes := []string{
		"image/jpeg",
		"image/jpg",
		"image/png",
		"image/gif",
		"image/webp",
		"image/svg+xml",
		"image/tiff",
		"image/bmp",
	}

	if strings.HasPrefix(vn.MimeType, "image/") {
		return true
	}

	if slices.Contains(imageMimeTypes, vn.MimeType) {
		return true
	}

	imageExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg", ".tiff", ".bmp"}
	ext := strings.ToLower(vn.Extension)
	return slices.Contains(imageExtensions, ext)
}

type Folder struct {
	ID       int64
	Name     string
	Path     string
	Posts    []*VaultNode
	Children map[string]*Folder
}

// Link represents a wikilink or markdown link
type Link struct {
	Target       string     `json:"target"` // the link target (filename or path)
	TargetID     int64      `json:"targetId"`
	TargetNode   *VaultNode `json:"targetNode"`
	Display      string     `json:"display"`                // display text if different from target
	Type         string     `json:"type"`                   // "wikilink" or "markdown"
	Line         int        `json:"line"`                   // line number where link appears
	IsResolved   bool       `json:"isResolved"`             // whether target file exists
	URL          string     `json:"url"`                    // url
	ResolvedPath string     `json:"resolvedPath,omitempty"` // actual path if resolved
}

type Footnote struct {
	ID     string
	Number int
	Text   string
}

type ExcalidrawPayload struct {
	Data string `json:"data"`
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

type Tag struct {
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Count int    `json:"count"`
}
