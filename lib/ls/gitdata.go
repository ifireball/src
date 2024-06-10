package ls

import (
	"time"

	"github.com/go-git/go-git/v5"
)

type RepoGitData struct {
	LastCommitTime time.Time
	MainRemoteURL string
}

func getRepoGitData(path string) (RepoGitData, error) {
	gitRepo, err := git.PlainOpen(path)
	if err != nil {
		return RepoGitData{}, err
	}
	lct, err := getLastCommitTime(gitRepo)
	if err != nil {
		return RepoGitData{}, err
	}
	mru, err := getMainRemoteURL(gitRepo)
	if err != nil {
		return RepoGitData{}, err
	}
	return RepoGitData{LastCommitTime: lct, MainRemoteURL: mru}, nil
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

func getMainRemoteURL(gitRepo *git.Repository) (string, error) {
	remotes, err := gitRepo.Remotes()
	if err != nil {
		return "" ,err
	}
	mainRemoteURL := ""
	urlRemoteName := ""
	for _, remote := range remotes {
		config := remote.Config()
		if len(config.URLs) <= 0 {
			continue
		}
		// Prefer the "upstream" remote most strongly, but use the first remote
		// we find otherwise
		if mainRemoteURL == "" || config.Name == "upstream" {
			mainRemoteURL = config.URLs[0]
			urlRemoteName = config.Name
			continue
		}
		// Prefer the "origin" remote to other remotes besides "upstream"
		if config.Name == "origin" && urlRemoteName != "upstream" {
			mainRemoteURL = config.URLs[0]
			urlRemoteName = config.Name
		}
	}
	return mainRemoteURL, nil
}
