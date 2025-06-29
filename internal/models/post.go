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

type BlogPost struct {
	Title        string
	FileName     string
	RawBody      []byte
	HTMLContent  template.HTML
	Date         time.Time
	Author       string
	UpdatedAt    *time.Time
	URL          string
	Images       []Image
	FilePath     string
	Tags         []Tag
	RelativePath string
	LinkedFrom   []Link
	LinkedTitles []string
}

type Image struct {
	RelativePath string
	Alt          string
	Width        string
	Height       string
}

type Link struct {
	Title string
	URL   string
}

type Tag struct {
	Name string
	Slug string
}

type Folder struct {
	Name     string
	Path     string
	Posts    []*BlogPost
	Children map[string]*Folder
}
