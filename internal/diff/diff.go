package diff

import (
	"crypto/sha256"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type File struct {
	Path string `toml:"path"`
	Hash string `toml:"hash"`
}

type LockFile struct {
	HashFormat string `toml:"hash_format"`
	Files      []File `toml:"files"`
}

type ChangeType int

const (
	Added ChangeType = iota
	Modified
	Deleted
)

func (c ChangeType) String() string {
	switch c {
	case Added:
		return "Added"
	case Modified:
		return "Modified"
	case Deleted:
		return "Deleted"
	default:
		return "Unknown"
	}
}

type Diff struct {
	File       string
	ChangeType ChangeType
}

type Tracker struct {
	directory       string
	outputDirectory string
	lockFile        string
}

func NewTracker(vaultDirectory, outputDirectory string) *Tracker {
	return &Tracker{
		directory:       vaultDirectory,
		outputDirectory: outputDirectory,
		lockFile:        filepath.Join(vaultDirectory, "lock.toml"),
	}
}

func (t *Tracker) GetChanges() ([]Diff, error) {
	currentLockFile, err := t.generateLockFile()
	if err != nil {
		return nil, fmt.Errorf("failed to generate current lock file: %w", err)
	}

	if _, err := os.Stat(t.lockFile); os.IsNotExist(err) {
		var diffs []Diff
		for _, file := range currentLockFile.Files {
			diffs = append(diffs, Diff{File: file.Path, ChangeType: Added})
		}
		return diffs, nil
	}

	oldLockFileData, err := os.ReadFile(t.lockFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read existing lock file: %w", err)
	}

	currentLockFileData, err := toml.Marshal(currentLockFile)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal current lock file: %w", err)
	}

	return getLockFilesChanged(oldLockFileData, currentLockFileData)
}

func (t *Tracker) UpdateLockFile() ([]Diff, error) {
	diffs, err := t.GetChanges()
	if err != nil {
		return nil, fmt.Errorf("failed to get changes: %w", err)
	}

	lockFile, err := t.generateLockFile()
	if err != nil {
		return nil, fmt.Errorf("failed to generate lock file: %w", err)
	}

	lockFileData, err := toml.Marshal(lockFile)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal lock file: %w", err)
	}

	if err := os.WriteFile(t.lockFile, lockFileData, 0644); err != nil {
		return nil, fmt.Errorf("failed to write lock file: %w", err)
	}

	return diffs, nil
}

func (t *Tracker) IsOutputDirectoryExists() bool {
	_, err := os.Stat(t.directory)
	return !os.IsNotExist(err)
}

func (t *Tracker) generateLockFile() (*LockFile, error) {
	var filePaths []string
	err := filepath.WalkDir(t.directory, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		filePaths = append(filePaths, path)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to walk directory: %w", err)
	}

	lockFile := &LockFile{
		HashFormat: "sha256",
		Files:      make([]File, 0, len(filePaths)),
	}

	for _, path := range filePaths {
		if filepath.Base(path) == "lock.toml" {
			continue
		}

		hash, err := t.calculateFileHash(path)
		if err != nil {
			return nil, fmt.Errorf("failed to calculate hash for %s: %w", path, err)
		}

		fileEntry := File{
			Path: path,
			Hash: hash,
		}

		lockFile.Files = append(lockFile.Files, fileEntry)
	}

	return lockFile, nil
}

func (t *Tracker) calculateFileHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func getLockFilesChanged(oldData, newData []byte) ([]Diff, error) {
	var oldLockFile LockFile
	if err := toml.Unmarshal(oldData, &oldLockFile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal old lock file: %w", err)
	}

	var newLockFile LockFile
	if err := toml.Unmarshal(newData, &newLockFile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal new lock file: %w", err)
	}

	oldFiles := make(map[string]string)
	newFiles := make(map[string]string)

	for _, file := range oldLockFile.Files {
		oldFiles[file.Path] = file.Hash
	}

	for _, file := range newLockFile.Files {
		newFiles[file.Path] = file.Hash
	}

	var diffs []Diff

	for path, newHash := range newFiles {
		if oldHash, exists := oldFiles[path]; exists {
			if oldHash != newHash {
				diffs = append(diffs, Diff{File: path, ChangeType: Modified})
			}
		} else {
			diffs = append(diffs, Diff{File: path, ChangeType: Added})
		}
	}

	for path := range oldFiles {
		if _, exists := newFiles[path]; !exists {
			diffs = append(diffs, Diff{File: path, ChangeType: Deleted})
		}
	}

	return diffs, nil
}

func (t *Tracker) HasChanges() (bool, error) {
	diffs, err := t.GetChanges()
	if err != nil {
		return false, err
	}
	return len(diffs) > 0, nil
}

func (t *Tracker) GetChangedFiles() ([]string, error) {
	diffs, err := t.GetChanges()
	if err != nil {
		return nil, err
	}

	var changedFiles []string
	for _, diff := range diffs {
		changedFiles = append(changedFiles, diff.File)
	}

	return changedFiles, nil
}

func (t *Tracker) GetChangedFilesByType() (map[ChangeType][]string, error) {
	diffs, err := t.GetChanges()
	if err != nil {
		return nil, err
	}

	result := make(map[ChangeType][]string)
	for _, diff := range diffs {
		result[diff.ChangeType] = append(result[diff.ChangeType], diff.File)
	}

	return result, nil
}
