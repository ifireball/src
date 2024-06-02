package ls

import (
	"errors"
	"fmt"
	"io/fs"
	"path/filepath"
)

type Repo struct {
	path string
}

func Repos(srcFs fs.FS) (<-chan Repo, error) {
	out := make(chan Repo)
	erc := make(chan error)

	go func() {
		defer func() { close(out) }()
		firstCall := true

		fs.WalkDir(srcFs, ".", func(path string, d fs.DirEntry, err error) error {
			defer func() { if firstCall { close(erc); firstCall = false } }()
			if err != nil {
				if d != nil {
					// Reading some nested dierectory failed, just skip it
					return fs.SkipDir
				}
				// Reading srcPath failed, so return the error
				erc <- err
				return err
			}
			if !d.IsDir() {
				return nil
			}
			git, err := srcFs.Open(filepath.Join(path, ".git"))
			if err != nil {
				if errors.Is(err, fs.ErrNotExist) {
					return nil
				}
				return err
			}
			defer func() { git.Close() }()
			stt, err := git.Stat()
			if err != nil {
				return err
			}
			if stt.IsDir() {
				out <- Repo{path: path}
				return fs.SkipDir
			}
			return nil
		})
	}()

	if err, ok := <- erc; ok {
		return nil, err
	}
	return out, nil
}

func Print(repos <-chan Repo) {
	for repo := range repos {
		fmt.Println(repo.path)
	}
}
