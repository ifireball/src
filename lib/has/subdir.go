package has

import (
	"errors"
	"io/fs"
	"path/filepath"
)

// Return wither a given path has a given sub directory
func SubDir(srcFs fs.FS, path, subDir string) (bool, error) {
	git, err := srcFs.Open(filepath.Join(path, subDir))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return false, nil
		}
		return false, err
	}
	defer func() { git.Close() }()
	stt, err := git.Stat()
	if err != nil {
		return false, err
	}
	if stt.IsDir() {
		return true, nil
	}
	return false, nil
}
