package transformers

import (
	"fmt"
	"gobsidian/internal/models"
)

type Transformer interface {
	Transform(content string, note *models.ParsedNote, ctx *TransformContext) (string, error)
	Name() string
}

type Pipeline struct {
	transformers []Transformer
}

func NewPipeline(transformers ...Transformer) *Pipeline {
	return &Pipeline{
		transformers: transformers,
	}
}

func (p *Pipeline) Transform(content string, note *models.ParsedNote, ctx *TransformContext) (string, error) {
	result := content
	var err error

	for _, transformer := range p.transformers {
		result, err = transformer.Transform(result, note, ctx)
		if err != nil {
			return "", fmt.Errorf("transformer %s failed: %w", transformer.Name(), err)
		}
	}

	return result, nil
}
