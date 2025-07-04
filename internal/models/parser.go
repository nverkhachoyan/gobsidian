package models

import "regexp"

type ParserRegexes struct {
	FrontmatterRegex   *regexp.Regexp
	WikilinkRegex      *regexp.Regexp
	ObsidianImageRegex *regexp.Regexp
	HashtagRegex       *regexp.Regexp
}
