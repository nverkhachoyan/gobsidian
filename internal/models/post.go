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
	Images       []string
	FilePath     string
	Tags         []Tag
	LinkedFrom   []Link
	LinkedTitles []string
}

type Link struct {
	Title string
	URL   string
}

type Tag struct {
	Name string
	Slug string
}
