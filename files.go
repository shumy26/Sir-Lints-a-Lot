package main

import (
	"os"
	"path/filepath"
	"strings"
)

//const python_files = "./python_files/"

func grabFiles(root string, files *[]string) error {
	dir, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, entry := range dir {
		fullPath := filepath.Join(root, entry.Name())
		if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
			if err := grabFiles(fullPath, files); err != nil {
				return err
			}
		} else if strings.HasSuffix(entry.Name(), ".py") {
			fullPath := filepath.Join(root, entry.Name())
			*files = append(*files, fullPath)
		}
	}
	return nil
}
