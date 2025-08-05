package crawler

import (
	"crypto/md5"
	"fmt"
	"gobsidian/internal/models"
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

func (vc *VaultCrawler) findInDirectory(candidates []*models.VaultNode, dirPath string) *models.VaultNode {
	for _, candidate := range candidates {
		if filepath.Dir(candidate.Path) == dirPath {
			return candidate
		}
	}
	return nil
}

func (vc *VaultCrawler) findInSubdirectories(candidates []*models.VaultNode, parentDir string) *models.VaultNode {
	for _, candidate := range candidates {
		candidateDir := filepath.Dir(candidate.Path)

		if vc.isSubdirectory(candidateDir, parentDir) {
			return candidate
		}
	}
	return nil
}

func (vc *VaultCrawler) findInVault(candidates []*models.VaultNode, excludeDir string) *models.VaultNode {
	// Sort candidates by depth (prefer files closer to root)
	sortedCandidates := make([]*models.VaultNode, len(candidates))
	copy(sortedCandidates, candidates)

	// Simple depth-based sorting (count path separators)
	for i := 0; i < len(sortedCandidates)-1; i++ {
		for j := i + 1; j < len(sortedCandidates); j++ {
			depthI := strings.Count(sortedCandidates[i].Path, string(filepath.Separator))
			depthJ := strings.Count(sortedCandidates[j].Path, string(filepath.Separator))

			if depthI > depthJ {
				sortedCandidates[i], sortedCandidates[j] = sortedCandidates[j], sortedCandidates[i]
			}
		}
	}

	for _, candidate := range sortedCandidates {
		candidateDir := filepath.Dir(candidate.Path)

		// Skip if we already checked this location
		if candidateDir == vc.RootPath ||
			candidateDir == excludeDir ||
			vc.isSubdirectory(candidateDir, excludeDir) {
			continue
		}

		return candidate
	}

	return nil
}

func (vc *VaultCrawler) findLineNumber(content, needle string, occurrence int) int {
	lines := strings.Split(content, "\n")
	found := 0

	for i, line := range lines {
		if strings.Contains(line, needle) {
			if found == occurrence {
				return i + 1
			}
			found++
		}
	}
	return 0
}

func (vc *VaultCrawler) normalizeDir(s string) string {
	s = strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})

	s = strings.ToLower(s)

	s = filepath.ToSlash(s)

	return s
}

func (vc *VaultCrawler) isSubdirectory(childPath, parentPath string) bool {
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
