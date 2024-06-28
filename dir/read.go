package dir

import (
	"os"
	"path/filepath"
)

type Dir struct {
	Path  string
	Files Files
}

type Files map[string]os.FileInfo

func ReadDir(path string) (Dir, error) {
	files, err := read(path)
	if err != nil {
		return Dir{}, err
	}

	return Dir{
		Path:  path,
		Files: files,
	}, nil
}

// read reads all files and directories in a given path.
func read(dir string) (Files, error) {
	fileMap := make(Files)
	return fileMap, filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			// should never happen
			panic(err)
		}
		filename, _ := filepath.Split(relPath)
		if filename != "." {
			fileMap[relPath] = info
		}
		return nil
	})
}
