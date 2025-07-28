package main

import (
	"os"
	"path/filepath"
	"strings"
)

type file struct {
	name string
	path string
}

func grabFiles(root string, files *[]file) error {
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
			var file file
			file.name = entry.Name()
			file.path = filepath.Join(root, entry.Name())
			*files = append(*files, file)
		}
	}
	return nil
}
