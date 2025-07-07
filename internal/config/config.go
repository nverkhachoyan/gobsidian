package config

import (
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"github.com/charmbracelet/log"
	"github.com/pelletier/go-toml/v2"
)

var frontmatterRegex = regexp.MustCompile(`(?s)^---\s*\n(.*?)\n---`)
var ObsidianImageRegex = regexp.MustCompile(`!\[\[([^|\]]+\.(?:png|jpg|jpeg|gif|bmp|svg))(?:\|(\d+)(?:x(\d+))?)?\]\]`)
var hashtagRegex = regexp.MustCompile(`\B#([a-zA-Z0-9-_]+)`)
var wikilinkRegex = regexp.MustCompile(`!?\[\[([^|\]]+)(?:\|([^\]]+))?\]\]`)

type Config struct {
	SiteConfig SiteConfig
	Regexes    Regexes
	Templates  *template.Template
	Logger     *log.Logger
}

type SiteConfig struct {
	SiteTitle       string `toml:"site_title"`
	SiteSubtitle    string `toml:"site_subtitle"`
	BaseURL         string `toml:"base_url"`
	NotesPerPage    int    `toml:"notes_per_page"`
	InputDirectory  string `toml:"input_directory"`
	OutputDirectory string `toml:"output_directory"`
	Env             string `toml:"env"`
}

type Regexes struct {
	ObsidianImageRegex *regexp.Regexp
	FrontmatterRegex   *regexp.Regexp
	HashtagRegex       *regexp.Regexp
	WikilinkRegex      *regexp.Regexp
}

func LoadConfig(filePath string) (Config, error) {
	var templates *template.Template

	templateDir := "templates"
	templates = template.New("main")

	err := filepath.WalkDir(templateDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && filepath.Ext(path) == ".html" {
			templates, err = templates.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to parse templates: %s", err)
	}

	yamlContent, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Failed to read config file: %s", err)
		log.Printf("Using default config")
		return Config{
			SiteConfig: SiteConfig{
				InputDirectory:  "content",
				OutputDirectory: "public",
				SiteTitle:       "My Blog",
				SiteSubtitle:    "A blog about my life",
				BaseURL:         "/",
			},
			Regexes: Regexes{
				ObsidianImageRegex: ObsidianImageRegex,
				FrontmatterRegex:   frontmatterRegex,
				HashtagRegex:       hashtagRegex,
				WikilinkRegex:      wikilinkRegex,
			},
			Templates: templates,
		}, nil
	}

	var cfg Config
	var siteConfig SiteConfig
	if err := toml.Unmarshal(yamlContent, &siteConfig); err != nil {
		return Config{}, err
	}

	cfg.SiteConfig = siteConfig

	cfg.Regexes = Regexes{
		ObsidianImageRegex: ObsidianImageRegex,
		FrontmatterRegex:   frontmatterRegex,
		HashtagRegex:       hashtagRegex,
		WikilinkRegex:      wikilinkRegex,
	}

	cfg.Templates = templates

	loggerLevel := log.InfoLevel
	if cfg.SiteConfig.Env == "development" {
		loggerLevel = log.DebugLevel
	}

	logger := log.NewWithOptions(os.Stderr, log.Options{
		TimeFormat:      time.Kitchen,
		ReportTimestamp: true,
		Level:           loggerLevel,
	})

	cfg.Logger = logger

	return cfg, nil
}
