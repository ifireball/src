package ls

import (
	"github.com/ifireball/src/lib/chu"
	"io/fs"
)

type Repo interface {
	RepoPathData
	RepoGitData
}

type repoImpl struct {
	Path string
	repoPathDataImpl
	repoGitDataImpl
}

func Repos(srcFs fs.FS, srcPath string) (<-chan Repo, error) {
	paths, err := repoPaths(srcFs, srcPath)
	if err != nil {
		return nil, err
	}
	repos := chu.Map(paths, func(path string) (Repo, bool) {
		rpd := getRepoPathData(srcPath, path)
		rgd, err := getRepoGitData(path)
		if err != nil {
			return nil, false
		}
		return &repoImpl{Path: path, repoPathDataImpl: *rpd, repoGitDataImpl: *rgd}, true
	})
	return repos, nil
}
