package ls

import (
	"io/fs"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/ifireball/src/lib/chu"
)

type Repo struct {
	Path string
	LastCommitTime time.Time
}

func Repos(srcFs fs.FS, srcPath string) (<-chan Repo, error) {
	paths, err := repoPaths(srcFs, srcPath)
	if err != nil {
		return nil, err
	}
	repos := chu.Map(paths, func(path string) (Repo, bool) {
		gitRepo, err := git.PlainOpen(path)
		if err != nil {
			return Repo{}, false
		}
		logIter, err := gitRepo.Log(&git.LogOptions{Order: git.LogOrderCommitterTime, All: true})
		if err != nil {
			return Repo{}, false
		}
		lastCommit, err := logIter.Next()
		logIter.Close()
		if err != nil {
			return Repo{}, false
		}
		return Repo{Path: path, LastCommitTime: lastCommit.Committer.When}, true
	})
	return repos, nil
}
