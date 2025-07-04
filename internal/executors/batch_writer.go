package executors

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/charmbracelet/log"
)

type FileToWrite struct {
	Path     string
	Content  []byte
	Template string
	Data     any
}

type BatchWriter struct {
	outputDir  string
	logger     *log.Logger
	templates  *template.Template
	bufferPool sync.Pool
	files      []FileToWrite
	mu         sync.Mutex
}

func NewBatchWriter(outputDir string, logger *log.Logger, templates *template.Template) *BatchWriter {
	return &BatchWriter{
		outputDir: outputDir,
		logger:    logger,
		templates: templates,
		bufferPool: sync.Pool{
			New: func() any {
				return &bytes.Buffer{}
			},
		},
	}
}

func (bw *BatchWriter) AddFile(relativePath, templateName string, data any) error {
	buf := bw.bufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	defer bw.bufferPool.Put(buf)

	if err := bw.templates.ExecuteTemplate(buf, templateName, data); err != nil {
		return fmt.Errorf("template execution failed for %s: %w", relativePath, err)
	}

	bw.mu.Lock()
	bw.files = append(bw.files, FileToWrite{
		Path:     filepath.Join(bw.outputDir, relativePath),
		Content:  append([]byte(nil), buf.Bytes()...),
		Data:     data,
		Template: templateName,
	})
	bw.mu.Unlock()

	return nil
}

func (bw *BatchWriter) WriteAllFiles() error {
	if err := bw.createDirectories(); err != nil {
		return err
	}

	return bw.writeFilesConcurrently()
}

func (bw *BatchWriter) createDirectories() error {
	dirs := make(map[string]bool)

	for _, file := range bw.files {
		dir := filepath.Dir(file.Path)
		dirs[dir] = true
	}

	for dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

func (bw *BatchWriter) writeFilesConcurrently() error {
	numWorkers := min(runtime.NumCPU(), len(bw.files))
	filesChan := make(chan FileToWrite, len(bw.files))
	errorsChan := make(chan error, len(bw.files))

	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for file := range filesChan {
				if err := os.WriteFile(file.Path, file.Content, 0644); err != nil {
					errorsChan <- fmt.Errorf("failed to write file %s: %w", file.Path, err)
				}
			}
		}()
	}

	for _, file := range bw.files {
		filesChan <- file
	}
	close(filesChan)

	wg.Wait()
	close(errorsChan)

	var errors []error
	for err := range errorsChan {
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		for _, err := range errors {
			bw.logger.Error("failed to write file", "error", err)
		}
		return fmt.Errorf("failed to write %d files", len(errors))
	}

	return nil
}
