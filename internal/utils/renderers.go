package utils

import (
	"regexp"
	"strings"
)

// Slugify converts a string into a URL-friendly "slug", preserving forward slashes
// and ensuring no hyphens are directly adjacent to slashes, and no multiple slashes.
func Slugify(s string) string {
	if s == "" {
		return ""
	}

	// 1. Convert to lowercase
	s = strings.ToLower(s)

	// 2. Replace any character that is NOT a lowercase letter, a digit, a forward slash, or a dot with a hyphen.
	// This also collapses multiple non-allowed characters into a single hyphen.
	// Example: "foo bar" -> "foo-bar", "foo--bar" -> "foo-bar", "foo  /  bar" -> "foo-/-bar"
	s = regexp.MustCompile(`[^a-z0-9/.]+`).ReplaceAllString(s, "-")

	// 3. Collapse any sequence of two or more hyphens into a single hyphen.
	// This ensures "foo--bar" becomes "foo-bar" if it wasn't already handled by step 2.
	s = regexp.MustCompile(`-{2,}`).ReplaceAllString(s, "-")

	// 4. Remove hyphens directly adjacent to slashes.
	//    a. Replace one or more hyphens immediately preceding a slash with just a slash.
	//       Example: "foo-/bar", "foo--/bar" -> "foo/bar"
	s = regexp.MustCompile(`-+/`).ReplaceAllString(s, "/")

	//    b. Replace one or more hyphens immediately following a slash with just a slash.
	//       Example: "foo/-bar", "foo/--bar" -> "foo/bar"
	s = regexp.MustCompile(`/-+`).ReplaceAllString(s, "/")

	// 5. Compress any sequence of two or more slashes into a single slash.
	// Example: "foo//bar" -> "foo/bar"
	s = regexp.MustCompile(`/{2,}`).ReplaceAllString(s, "/")

	// 6. Trim leading and trailing hyphens and slashes from the entire string.
	// This is important to catch cases like "/-foo" or "foo-/" at the very beginning or end
	// after all other replacements.
	s = strings.Trim(s, "-/")

	return s
}

func Deslugify(s string) string {
	return strings.ReplaceAll(s, "-", " ")
}
