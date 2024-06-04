package ls

import (
	"time"

	"github.com/go-git/go-git/v5"
)

type RepoGitData struct {
	LastCommitTime time.Time
}

func getRepoGitData(path string) (RepoGitData, error) {
	gitRepo, err := git.PlainOpen(path)
	if err != nil {
		return RepoGitData{}, err
	}
	lct, err := getLastCommitTime(gitRepo)
	return RepoGitData{LastCommitTime: lct}, nil
}

func getLastCommitTime(gitRepo *git.Repository) (time.Time, error) {
	logIter, err := gitRepo.Log(&git.LogOptions{Order: git.LogOrderCommitterTime, All: true})
	if err != nil {
		return time.Time{}, err
	}
	lastCommit, err := logIter.Next()
	logIter.Close()
	if err != nil {
		return time.Time{}, err
	}
	return lastCommit.Committer.When, nil
}
