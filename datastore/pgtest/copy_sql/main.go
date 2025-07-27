package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func run() error {
	sourceDir := "../../meta/dev/sql"
	destDir := "sql"

	if err := os.MkdirAll(destDir, 0755); err != nil {
		return fmt.Errorf("creating destination directory: %w", err)
	}

	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".sql" {
			return nil
		}

		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return fmt.Errorf("calculating relative path: %w", err)
		}

		destPath := filepath.Join(destDir, relPath)

		if err := copyFile(path, destPath); err != nil {
			return fmt.Errorf("copying file %s to %s: %w", path, destPath, err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("walking source directory: %w", err)
	}

	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("opening source file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("creating destination file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("copying file contents: %w", err)
	}

	return nil
}
