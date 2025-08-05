package transformers

import (
	"encoding/json"
	"fmt"
	"gobsidian/internal/models"
	"mime"
	"path/filepath"
	"strconv"

	lzstring "github.com/daku10/go-lz-string"
)

type ExcalidrawEnricher struct {
}

func NewExcalidrawEnricher() *ExcalidrawEnricher {
	return &ExcalidrawEnricher{}
}

func (ip *ExcalidrawEnricher) Name() string {
	return "excalidraw-enricher"
}

func (he *ExcalidrawEnricher) Transform(content string, note *models.VaultNode, ctx *TransformContext) (string, error) {
	if note.NoteType != models.NoteTypeExcalidraw {
		return content, nil
	}

	initialData := note.Excalidraw.InitialData

	for _, element := range initialData.Elements {
		if element.Type == models.ElementTypeText {
			if text, ok := note.Excalidraw.TextElements[element.ID]; ok {
				element.Text = text
			}
		}
	}

	// Hydrate file elements
	for fileID, file := range note.Excalidraw.EmbeddedFiles {
		mimeType := mime.TypeByExtension(filepath.Ext(file))
		initialData.Files[fileID] = &models.ExcalidrawFile{
			ID:       fileID,
			DataURL:  "http://" + ctx.BaseURL + ":" + strconv.Itoa(ctx.Port) + "/images/" + fileID + filepath.Ext(file),
			MimeType: mimeType,
		}
	}

	initialDataJSON, err := json.Marshal(initialData)
	if err != nil {
		fmt.Printf("failed to marshal initial data: %v\n", err)
		return content, fmt.Errorf("failed to marshal initial data: %w", err)
	}

	initialDataBase64, err := lzstring.CompressToBase64(string(initialDataJSON))
	if err != nil {
		fmt.Printf("failed to compress initial data: %v\n", err)
		return content, fmt.Errorf("failed to compress initial data: %w", err)
	}

	note.ExcalidrawPayload = models.ExcalidrawPayload{
		Data: initialDataBase64,
	}

	return content, nil
}
