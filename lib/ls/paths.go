package ls

import (
	"io/fs"
	
	"github.com/ifireball/src/lib/has"
)

// Return a channel of all the directories in the given FS that contain a .git
// subdirectory, while avoiding further descending into such directories.
func repoPaths(srcFs fs.FS) (<-chan string, error) {
	out := make(chan string)
	erc := make(chan error)

	go func() {
		defer func() { close(out) }()
		firstCall := true

		fs.WalkDir(srcFs, ".", func(path string, d fs.DirEntry, err error) error {
			defer func() { if firstCall { close(erc); firstCall = false } }()
			if err != nil {
				if d != nil {
					// Reading some nested directory failed, just skip it
					return fs.SkipDir
				}
				// Reading srcPath failed, so return the error
				erc <- err
				return err
			}
			if !d.IsDir() {
				return nil
			}
			if hasGit, err := has.SubDir(srcFs, path, ".git"); err == nil && hasGit {
				out <- path
				return fs.SkipDir
			} else {
				return err
			}
		})
	}()

	if err, ok := <- erc; ok {
		return nil, err
	}
	return out, nil
}
