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
	ID                      int64
	Title                   string
	FileName                string
	RawFileName             string
	RawBody                 []byte
	HTMLContent             template.HTML
	Date                    time.Time
	Author                  string
	UpdatedAt               *time.Time
	URL                     string
	Images                  []Image
	FilePath                string
	Tags                    []Tag
	RelativePath            string
	RelativePathWithoutName string
	LinkedFrom              []Link
	LinkedTitlesStrings     []string
	LinkedTitles            []*BlogPost
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
	Posts    []*BlogPost
	Children map[string]*Folder
}
