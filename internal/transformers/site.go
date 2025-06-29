package transformers

import (
	"fmt"
	"html/template"
	"sync"

	"github.com/charmbracelet/log"

	"os"
	"path/filepath"
	"sort"
	"strings"

	"gobsidian/internal/config"
	"gobsidian/internal/executors"
	"gobsidian/internal/models"
	"gobsidian/internal/parsers"
	"gobsidian/internal/utils"

	"github.com/gomarkdown/markdown/html"

	"github.com/gomarkdown/markdown"
)

type SiteTransformer struct {
	Config           config.Config
	Logger           *log.Logger
	Templates        *template.Template
	Parser           parsers.Parser
	MarkdownRenderer *html.Renderer
}

func NewSiteTransformer(
	config config.Config,
	logger *log.Logger,
	parser parsers.Parser,
	markdownRenderer *html.Renderer,

) (*SiteTransformer, error) {
	templates, err := template.ParseGlob("templates/*.html")
	if err != nil {
		return nil, fmt.Errorf("failed to parse templates: %w", err)
	}
	return &SiteTransformer{
		Config:           config,
		Logger:           logger,
		Templates:        templates,
		Parser:           parser,
		MarkdownRenderer: markdownRenderer,
	}, nil
}

func (g *SiteTransformer) GenerateSite() error {
	g.Logger.Print("Generating site...")

	// --- Step 1: Create output dirs ---
	if err := g.createOutputDirectories(); err != nil {
		return fmt.Errorf("failed to create directories: %w", err)
	}

	// --- Step 2: Discover and parse notes ---
	posts, titleToPost, err := g.discoverAndParseNotes()
	if err != nil {
		return fmt.Errorf("error during initial walk: %w", err)
	}

	// --- Step 3: Apply transformations ---
	for _, post := range posts {
		postBody := string((*post).RawBody)

		contentWithImagesReplaced := g.replaceObsidianImageLinks(postBody)
		contentWithWikilinks := g.replaceWikilinksAndResolveBacklinks(contentWithImagesReplaced, post, titleToPost)
		hashtagAsLinks := g.enrichHashtagsWithLinks(contentWithWikilinks)

		(*post).HTMLContent = template.HTML(markdown.ToHTML([]byte(hashtagAsLinks), nil, g.MarkdownRenderer))
	}

	// --- Step 4: Create subdirectories ---
	fileTree, numberOfFolders, err := g.buildFolderTreeAndCreateDirs(posts)
	if err != nil {
		return fmt.Errorf("failed to create subdirectories: %w", err)
	}

	// --- Step 5: Execute the templates ---
	numberOfTags, err := g.executeTemplates(posts, fileTree)
	if err != nil {
		return err
	}

	// --- Step 6: Copy Assets ---
	if err := g.copyAssets(posts); err != nil {
		g.Logger.Warn("Failed to copy some images", "error", err)
	}

	g.Logger.Print("Site generation complete!",
		"posts", len(posts),
		"tags", numberOfTags,
		"index_pages", 1,
		"preview_pages", len(posts),
		"folders", numberOfFolders,
	)
	return nil
}

func (g *SiteTransformer) replaceObsidianImageLinks(body string) string {
	return g.Config.RegexpConfig.ObsidianImageRegex.ReplaceAllStringFunc(string(body), func(match string) string {
		parts := g.Config.RegexpConfig.ObsidianImageRegex.FindStringSubmatch(match)
		if len(parts) > 1 {
			imageName := parts[1]
			return fmt.Sprintf(`<img src="/images/%s" alt="%s" height="200" />`, imageName, imageName)
		}
		return match
	})
}

func (g *SiteTransformer) buildFolderTreeAndCreateDirs(posts []*models.BlogPost) (*models.Folder, int, error) {
	root := &models.Folder{
		Name:     "Home",
		Path:     "",
		Posts:    []*models.BlogPost{},
		Children: make(map[string]*models.Folder),
	}
	numberOfFolders := 0

	for _, post := range posts {
		if post.RelativePath == "" || post.RelativePath == "." {
			root.Posts = append(root.Posts, post)
			continue
		}

		parts := strings.Split(post.RelativePath, string(filepath.Separator))
		currentNode := root

		for _, part := range parts {
			if part == "" {
				continue
			}

			if _, ok := currentNode.Children[part]; !ok {
				childPath := filepath.ToSlash(filepath.Join(currentNode.Path, part))

				postDir := filepath.Join(g.Config.OutputDirectory, childPath)
				if err := os.MkdirAll(postDir, 0755); err != nil {
					return nil, 0, fmt.Errorf("failed to create post subdirectory %s: %w", postDir, err)
				}
				previewDir := filepath.Join(g.Config.OutputDirectory, "previews", childPath)
				if err := os.MkdirAll(previewDir, 0755); err != nil {
					return nil, 0, fmt.Errorf("failed to create preview subdirectory %s: %w", previewDir, err)
				}

				currentNode.Children[part] = &models.Folder{
					Name:     part,
					Path:     childPath,
					Posts:    []*models.BlogPost{},
					Children: make(map[string]*models.Folder),
				}
				numberOfFolders++
			}
			currentNode = currentNode.Children[part]
		}

		currentNode.Posts = append(currentNode.Posts, post)
	}

	sortFolderTree(root)
	return root, numberOfFolders, nil
}

func sortFolderTree(folder *models.Folder) {
	// Sort posts in the current folder alphabetically by title
	sort.Slice(folder.Posts, func(i, j int) bool {
		return folder.Posts[i].Title < folder.Posts[j].Title
	})

	// Recurse into children
	for _, child := range folder.Children {
		sortFolderTree(child)
	}
}

func (g *SiteTransformer) replaceWikilinksAndResolveBacklinks(body string, pd *models.BlogPost, titleToPost map[string]*models.BlogPost) string {
	return g.Config.RegexpConfig.WikilinkRegex.ReplaceAllStringFunc(body, func(match string) string {
		parts := g.Config.RegexpConfig.WikilinkRegex.FindStringSubmatch(match)
		linkTarget := strings.TrimSpace(parts[1])
		linkText := linkTarget
		if len(parts) > 2 && parts[2] != "" {
			linkText = parts[2]
		}

		if linkedPd, ok := titleToPost[linkTarget]; ok {
			// Add backlink to the target post
			// Avoid duplicates
			found := false
			for _, l := range linkedPd.LinkedFrom {
				if l.URL == pd.URL {
					found = true
					break
				}
			}
			if !found {
				linkedPd.LinkedFrom = append(linkedPd.LinkedFrom, models.Link{
					Title: pd.Title,
					URL:   pd.URL,
				})
			}
			return fmt.Sprintf(`<a href="%s">%s</a>`, linkedPd.URL, linkText)
		}
		// If link target doesn't exist, return plain text
		return linkText
	})
}

func (g *SiteTransformer) enrichHashtagsWithLinks(body string) string {
	return g.Config.RegexpConfig.HashtagRegex.ReplaceAllStringFunc(body, func(match string) string {
		parts := g.Config.RegexpConfig.HashtagRegex.FindStringSubmatch(match)
		tagName := parts[1]
		return fmt.Sprintf(`<a class="tag" href="/tag/%s">#%s</a>`, tagName, tagName)
	})
}

func (g *SiteTransformer) executeTemplates(blogPosts []*models.BlogPost, fileTree *models.Folder) (int, error) {
	tagsMap := make(map[string]*models.Tag)
	for _, pd := range blogPosts {
		for _, tag := range pd.Tags {
			tagsMap[tag.Slug] = &tag
		}
	}

	tags := make([]models.Tag, 0)
	for _, tag := range tagsMap {
		tags = append(tags, models.Tag{Name: tag.Name, Slug: tag.Slug})
	}

	sort.Slice(tags, func(i, j int) bool {
		return tags[i].Name < tags[j].Name
	})

	for _, p := range blogPosts {
		if err := executors.ExecutePostPage(g.Config, *p); err != nil {
			g.Logger.Error("Error generating post page", "title", p.Title, "error", err)
		}
		if err := executors.ExecutePreviewPage(g.Config, *p); err != nil {
			g.Logger.Error("Error generating preview page", "title", p.Title, "error", err)
		}
	}

	postsByTag := make(map[string][]*models.BlogPost, len(tags))

	for _, p := range blogPosts {
		for _, tag := range p.Tags {
			postsByTag[tag.Slug] = append(postsByTag[tag.Slug], p)
		}
	}

	for _, tag := range tags {
		if err := executors.ExecuteTagPage(g.Config, tag, postsByTag[tag.Slug]); err != nil {
			g.Logger.Error("Error generating tag page", "name", tag.Name, "error", err)

		}
	}

	// This is new: recursively generate folder pages
	if err := g.executeFolderTemplates(fileTree, tags); err != nil {
		g.Logger.Error("Error generating folder pages", "error", err)
	}

	if err := executors.ExecuteIndexPage(g.Config, blogPosts, tags, fileTree); err != nil {
		return 0, err
	}

	if err := executors.ExecuteNotFoundPage(g.Config); err != nil {
		g.Logger.Error("Error generating not found page", "error", err)
	}

	return len(tags), nil
}

func (g *SiteTransformer) executeFolderTemplates(folder *models.Folder, allTags []models.Tag) error {
	if err := executors.ExecuteFolderPage(g.Config, folder, allTags); err != nil {
		return err
	}

	for _, child := range folder.Children {
		if err := g.executeFolderTemplates(child, allTags); err != nil {
			return err
		}
	}
	return nil
}

func (g *SiteTransformer) createOutputDirectories() error {
	if err := os.MkdirAll(filepath.Join(g.Config.OutputDirectory, "images"), 0755); err != nil {
		return fmt.Errorf("failed to create images directory %s: %w", filepath.Join(g.Config.OutputDirectory, "images"), err)
	}
	if err := os.MkdirAll(filepath.Join(g.Config.OutputDirectory, "previews"), 0755); err != nil {
		return fmt.Errorf("failed to create previews directory %s: %w", filepath.Join(g.Config.OutputDirectory, "previews"), err)
	}
	return nil
}

func (g *SiteTransformer) discoverAndParseNotes() ([]*models.BlogPost, map[string]*models.BlogPost, error) {
	var posts []*models.BlogPost
	titleToPost := make(map[string]*models.BlogPost)

	err := filepath.Walk(g.Config.InputDirectory, func(path string, info os.FileInfo, err error) error {

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			post, images, linkedTitles, err := g.Parser.ParseNote(path)
			if err != nil {
				g.Logger.Warn("Skipping file during initial parse", "path", path, "error", err)
				return nil
			}

			post.URL = post.RelativePath + "/" + post.FileName

			post.Images = images
			post.FilePath = path
			post.LinkedTitles = linkedTitles

			posts = append(posts, &post)
			titleToPost[post.Title] = &post
		}
		return nil
	})
	if err != nil {
		return nil, nil, fmt.Errorf("error during initial walk: %w", err)
	}

	return posts, titleToPost, nil
}

func (g *SiteTransformer) copyAssets(posts []*models.BlogPost) error {
	var wg sync.WaitGroup
	errorChan := make(chan error, 100)

	images := make(map[string]string)
	for _, p := range posts {
		for _, img := range p.Images {
			imagePath := filepath.Join(filepath.Dir(p.FilePath), img)
			images[img] = imagePath
		}
	}

	g.Logger.Print("Copying images 🖼️ ", "type", "images", "count", len(images))
	for name, path := range images {
		wg.Add(1)
		go func(name, path string) {
			defer wg.Done()
			destPath := filepath.Join(g.Config.OutputDirectory, "images", name)
			if err := utils.CopyFile(path, destPath); err != nil {
				errorChan <- fmt.Errorf("image %s: %w", name, err)
			}
		}(name, path)
	}

	for _, assetType := range []string{"js", "css"} {
		wg.Add(1)
		go func(assetType string) {
			defer wg.Done()
			g.Logger.Print("Copying static assets 📦", "type", assetType)
			if err := utils.CopyStaticDirectory(assetType, g.Config.OutputDirectory); err != nil {
				errorChan <- fmt.Errorf("asset folder '%s': %w", assetType, err)
			}
		}(assetType)
	}

	wg.Wait()
	close(errorChan)

	var allErrors []string
	for err := range errorChan {
		allErrors = append(allErrors, err.Error())
	}

	if len(allErrors) > 0 {
		return fmt.Errorf("failed to copy %d assets:\n- %s", len(allErrors), strings.Join(allErrors, "\n- "))
	}

	return nil
}
