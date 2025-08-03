package generators

import (
	"bytes"
	"fmt"
	"gobsidian/internal/crawler"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/alecthomas/chroma/v2/styles"

	"gobsidian/internal/utils"
	"regexp"
)

func (g *StaticSiteGenerator) copyImages(nodesIndex map[string]*crawler.VaultNode) time.Duration {
	start := time.Now()
	var wg sync.WaitGroup
	errorChan := make(chan error, 100)

	for _, node := range nodesIndex {
		wg.Add(1)
		go func(name, path string) {
			defer wg.Done()
			destPath := filepath.Join(g.SiteConfig.OutputDirectory, imagesDir, name)
			if err := utils.CopyFile(path, destPath); err != nil {
				errorChan <- fmt.Errorf("image %s: %w", name, err)

			}
		}(node.Name, node.Path)
	}

	wg.Wait()
	close(errorChan)

	for err := range errorChan {
		g.Logger.Debugf("%s\n", err.Error())
	}

	return time.Since(start)
}

func (g *StaticSiteGenerator) copyAssetsDir() time.Duration {
	start := time.Now()
	var wg sync.WaitGroup
	errorChan := make(chan error, 100)

	for _, dir := range []string{assetsDir, generatedDir, srcDir} {
		wg.Add(1)
		go func(dir string) {
			defer wg.Done()
			g.Logger.Debug("Copying static assets 📦", "dir", dir)
			if err := utils.CopyStaticDirectory(dir, g.SiteConfig.OutputDirectory); err != nil {
				errorChan <- fmt.Errorf("asset folder '%s': %w", dir, err)
			}
		}(dir)
	}

	wg.Wait()
	close(errorChan)

	_ = g.copyCssSnippets()

	for err := range errorChan {
		g.Logger.Debugf("%s\n", err.Error())
	}

	return time.Since(start)
}

func (g *StaticSiteGenerator) exportFiletreeJSON(fileTreeJSON []byte) time.Duration {
	start := time.Now()

	if len(fileTreeJSON) > 0 {

		destPath := filepath.Join(g.SiteConfig.OutputDirectory, generatedDir, fileTreeFilename)
		if err := os.MkdirAll(filepath.Dir(destPath), 0755); err != nil {
			g.Logger.Debugf("%s\n", err.Error())
		}
		if err := os.WriteFile(destPath, fileTreeJSON, 0644); err != nil {
			g.Logger.Debugf("%s\n", err.Error())
		}
	}

	return time.Since(start)
}

func (g *StaticSiteGenerator) copyCssSnippets() []string {
	var snippetPaths []string

	filepath.Walk(g.SiteConfig.InputDirectory+"/"+g.SiteConfig.SnippetsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), ".css") {
			snippetPaths = append(snippetPaths, path)
		}

		return nil
	})

	// TODO: Append CSS snippets to the CSS file
	// for _, snippet := range snippetPaths {
	// 	snippetText, err := os.ReadFile(snippet)
	// 	if err != nil {
	// 		g.Logger.Debugf("failed to read CSS snippet: %w", err)
	// 		continue
	// 	}

	// 	fmt.Println(string(snippetText))
	// }

	return snippetPaths
}

func (g *StaticSiteGenerator) generateSyntaxHighlighterCSS() error {
	var lightBuf, darkBuf bytes.Buffer
	formatter := chromahtml.New(chromahtml.WithClasses(true))

	if err := formatter.WriteCSS(
		&lightBuf,
		styles.Get(g.SiteConfig.Theme.SyntaxHighlighterLight),
	); err != nil {
		return fmt.Errorf("failed to generate light theme: %w", err)
	}

	if err := formatter.WriteCSS(
		&darkBuf,
		styles.Get(g.SiteConfig.Theme.SyntaxHighlighterDark),
	); err != nil {
		return fmt.Errorf("failed to generate dark theme: %w", err)
	}

	scopedDarkCSS := g.scopeCSS(darkBuf.String(), "html.dark-mode")

	// combine light and dark themes
	finalCSS := lightBuf.String() + "\n" + scopedDarkCSS
	outputPath := filepath.Join(g.SiteConfig.OutputDirectory, generatedDir, chromaCSSPath)

	if err := os.WriteFile(outputPath, []byte(finalCSS), 0644); err != nil {
		return fmt.Errorf("failed to write combined CSS file: %w", err)
	}

	g.Logger.Debug("Successfully generated combined Chroma CSS at %s", outputPath)
	return nil
}

// util function to scope CSS to a specific selector
func (g *StaticSiteGenerator) scopeCSS(css, scope string) string {
	var result strings.Builder
	r := regexp.MustCompile(`(?s)([^}]+)({[^}]+})`)
	submatches := r.FindAllStringSubmatch(css, -1)

	for _, submatch := range submatches {
		selectors := strings.TrimSpace(submatch[1])
		declarations := submatch[2]

		if strings.HasPrefix(selectors, "@") {
			result.WriteString(selectors)
			result.WriteString(declarations)
			continue
		}

		var scopedSelectors []string
		for _, selector := range strings.Split(selectors, ",") {
			scopedSelectors = append(scopedSelectors, scope+" "+strings.TrimSpace(selector))
		}
		result.WriteString(strings.Join(scopedSelectors, ", "))
		result.WriteString(declarations)
	}
	return result.String()
}
