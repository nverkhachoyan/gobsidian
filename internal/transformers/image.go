package transformers

import (
	"fmt"
	"gobsidian/internal/models"
	"regexp"
)

type ImageProcessor struct {
	imageLinkRegex *regexp.Regexp
}

func NewImageProcessor(regex *regexp.Regexp) *ImageProcessor {
	return &ImageProcessor{
		imageLinkRegex: regex,
	}
}

func (ip *ImageProcessor) Name() string {
	return "image-processor"
}

func (ip *ImageProcessor) Transform(content string, note *models.ParsedNote, ctx *TransformContext) (string, error) {
	result := content
	return ip.imageLinkRegex.ReplaceAllStringFunc(result, func(match string) string {
		parts := ip.imageLinkRegex.FindStringSubmatch(match)
		if len(parts) > 3 {
			imageName := parts[1]
			width := parts[2]
			height := parts[3]
			return fmt.Sprintf(`<img src="/images/%s" alt="%s" width="%s" height="%s" />`, imageName, imageName, width, height)
		}
		if len(parts) > 2 {
			imageName := parts[1]
			width := parts[2]
			return fmt.Sprintf(`<img src="/images/%s" alt="%s"  width="%s" />`, imageName, imageName, width)
		}
		if len(parts) > 1 {
			imageName := parts[1]
			return fmt.Sprintf(`<img src="/images/%s" alt="%s"  />`, imageName, imageName)
		}

		return match
	}), nil

}
