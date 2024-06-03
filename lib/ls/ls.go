package ls

import (
	"io/fs"

	"github.com/ifireball/src/lib/chu"
)

type Repo struct {
	path string
}

func Repos(srcFs fs.FS) (<-chan Repo, error) {
	paths, err := repoPaths(srcFs)
	if err != nil {
		return nil, err
	}
	repos := chu.Map(paths, func(path string) Repo {
		return Repo{path: path}
	})
	return repos, nil
}
