package dir

import (
	"io/fs"
	"path/filepath"
)

func FilesFS(fsys fs.FS, dir string) ([]string, error) {
	if dir == "" {
		dir = "."
	}
	var fileNames []string
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Check if it's a regular file
		if d.Type().IsRegular() {
			// Print the file name
			fileNames = append(fileNames, d.Name())
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return fileNames, nil
}
