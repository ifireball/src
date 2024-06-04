package ls

import (
	"io/fs"
	"github.com/ifireball/src/lib/chu"
)

type Repo struct {
	Path string
	RepoGitData
}

func Repos(srcFs fs.FS, srcPath string) (<-chan Repo, error) {
	paths, err := repoPaths(srcFs, srcPath)
	if err != nil {
		return nil, err
	}
	repos := chu.Map(paths, func(path string) (Repo, bool) {
		rgd, err := getRepoGitData(path)
		if err != nil {
			return Repo{}, false
		}
		return Repo{Path: path, RepoGitData: rgd}, true
	})
	return repos, nil
}
