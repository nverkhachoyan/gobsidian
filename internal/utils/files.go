package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("could not open source file %s: %w\n", src, err)
	}
	defer sourceFile.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return fmt.Errorf("could not create destination directory for %s: %w", dst, err)
	}

	destFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("could not create destination file %s: %w", dst, err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("could not copy from %s to %s: %w", src, dst, err)
	}
	return nil
}

func CopyStaticDirectory(sourceDir, destDir string) error {
	// It's not an error if the source directory doesn't exist.
	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		return nil
	}

	return filepath.Walk(sourceDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceDir, srcPath)
		if err != nil {
			return fmt.Errorf("could not get relative path for %s: %w", srcPath, err)
		}
		destPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		return CopyFile(srcPath, destPath)
	})
}

func IsDirectoryExists(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err == nil {
		return stat.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
