package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
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

	baseDataSrc := path.Join(sourceDir, "base_data.sql")
	schamaSrc := path.Join(sourceDir, "schema_relational.sql")
	baseDataDest := path.Join(destDir, "base_data.sql")
	schamaDest := path.Join(destDir, "schema_relational.sql")

	if err := copyFile(baseDataSrc, baseDataDest); err != nil {
		return fmt.Errorf("copying file %s to %s: %w", baseDataSrc, destDir, err)
	}

	if err := copyFile(schamaSrc, schamaDest); err != nil {
		return fmt.Errorf("copying file %s to %s: %w", schamaSrc, destDir, err)
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
