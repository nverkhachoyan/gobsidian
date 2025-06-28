package config

import (
	"html/template"
	"log"
	"os"
	"regexp"

	"github.com/goccy/go-yaml"
)

var frontmatterRegex = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---`)
var ObsidianImageRegex = regexp.MustCompile(`!\[\[([^|\]]+\.(?:png|jpg|jpeg|gif|bmp|svg))\]\]`)
var hashtagRegex = regexp.MustCompile(`\B#([a-zA-Z0-9-_]+)`)
var wikilinkRegex = regexp.MustCompile(`\[\[([^|\]]+)(?:\|([^\]]+))?\]\]`)

type Config struct {
	InputDirectory  string
	OutputDirectory string
	SiteTitle       string
	SiteSubtitle    string
	BaseURL         string
	RegexpConfig    RegexpConfig
	Templates       *template.Template
}

type RegexpConfig struct {
	ObsidianImageRegex *regexp.Regexp
	FrontmatterRegex   *regexp.Regexp
	HashtagRegex       *regexp.Regexp
	WikilinkRegex      *regexp.Regexp
}

func ReadConfig(filePath string) (Config, error) {
	var templates, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatalf("Failed to parse templates: %s", err)
	}

	yamlContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read config file: %s", err)
		log.Printf("Using default config")
		return Config{
			InputDirectory:  "content",
			OutputDirectory: "public",
			SiteTitle:       "My Blog",
			SiteSubtitle:    "A blog about my life",
			BaseURL:         "/",
			RegexpConfig: RegexpConfig{
				ObsidianImageRegex: ObsidianImageRegex,
				FrontmatterRegex:   frontmatterRegex,
				HashtagRegex:       hashtagRegex,
				WikilinkRegex:      wikilinkRegex,
			},
			Templates: templates,
		}, nil
	}

	var configYaml yaml.MapSlice
	if err := yaml.Unmarshal(yamlContent, &configYaml); err != nil {
		return Config{}, err
	}

	cfg := Config{}

	for _, v := range configYaml {
		switch v.Key {
		case "InputDirectory":
			cfg.InputDirectory = v.Value.(string)
		case "OutputDirectory":
			cfg.OutputDirectory = v.Value.(string)
		case "SiteTitle":
			cfg.SiteTitle = v.Value.(string)
		case "SiteSubtitle":
			cfg.SiteSubtitle = v.Value.(string)
		case "BaseURL":
			cfg.BaseURL = v.Value.(string)
		}
	}

	cfg.RegexpConfig = RegexpConfig{
		ObsidianImageRegex: ObsidianImageRegex,
		FrontmatterRegex:   frontmatterRegex,
		HashtagRegex:       hashtagRegex,
		WikilinkRegex:      wikilinkRegex,
	}

	cfg.Templates = templates

	return cfg, nil
}
