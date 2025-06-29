package utils

import (
	"regexp"
	"strings"
)

// Slugify converts a string into a URL-friendly "slug"
func Slugify(s string) string {
	if s == "" {
		return ""
	}
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}
