package models

type File struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type FolderNode struct {
	Name     string        `json:"name"`
	Path     string        `json:"path"`
	Children []*FolderNode `json:"children"`
	Files    []*File       `json:"files"`
}
