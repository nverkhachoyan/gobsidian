package crawler

import (
	"crypto/md5"
	"fmt"
	"hash/fnv"
	"io"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func generateIDFromPath(path string) int64 {
	h := fnv.New64a()
	h.Write([]byte(path))
	return int64(h.Sum64())
}

func (vc *VaultCrawler) normalizeDir(s string) string {
	s = strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	s = strings.ToLower(s)

	s = filepath.ToSlash(s)

	return s
}

// isSubdirectory checks if childPath is a subdirectory of parentPath
func (vc *VaultCrawler) isSubdirectory(childPath, parentPath string) bool {
	// Normalize paths
	childPath = filepath.Clean(childPath)
	parentPath = filepath.Clean(parentPath)

	// Make relative to check containment
	relPath, err := filepath.Rel(parentPath, childPath)
	if err != nil {
		return false
	}

	// If relative path starts with "..", it's not a subdirectory
	return !strings.HasPrefix(relPath, "..") && relPath != "."
}

func (vc *VaultCrawler) calculateHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
