package ls

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type RepoGitData interface {
	LastCommitTime() time.Time
	MainRemoteURL() string
	Config() *config.Config
}

type repoGitDataImpl struct {
	lastCommitTime time.Time
	mainRemoteURL string
	config *config.Config
}

func (rgd *repoGitDataImpl) LastCommitTime() time.Time {
	return rgd.lastCommitTime
}

func (rgd *repoGitDataImpl) MainRemoteURL() string {
	return rgd.mainRemoteURL
}

func (rgd *repoGitDataImpl) Config() *config.Config {
	return rgd.config
}

func getRepoGitData(path string) (*repoGitDataImpl, error) {
	gitRepo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}
	lct, err := getLastCommitTime(gitRepo)
	if err != nil {
		return nil, err
	}
	mru, err := getMainRemoteURL(gitRepo)
	if err != nil {
		return nil, err
	}
	cfg, err := gitRepo.Config()
	if err != nil {
		return nil, err
	}
	return &repoGitDataImpl{
		lastCommitTime: lct,
		mainRemoteURL: mru,
		config: cfg,
	}, nil
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
