package ls

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/ifireball/src/lib/taxonomy"
)

type RepoGitData interface {
	LastCommitTime() time.Time
	MainRemoteURL() string
	Config() *config.Config
	Branches() []taxonomy.Branch
	Head() string
}

type repoGitDataImpl struct {
	lastCommitTime time.Time
	mainRemoteURL string
	config *config.Config
	branches []taxonomy.Branch
	head string
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

func (rgd *repoGitDataImpl) Branches() []taxonomy.Branch {
	return rgd.branches
}

func (rgd *repoGitDataImpl) Head() string {
	return rgd.head
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
	branches, err := getBranches(gitRepo)
	if err != nil {
		return nil, err
	}
	head, err := gitRepo.Reference(plumbing.HEAD, false)
	if err != nil {
		return nil, err
	}
	return &repoGitDataImpl{
		lastCommitTime: lct,
		mainRemoteURL: mru,
		config: cfg,
		branches: branches,
		head: head.Strings()[1],
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

func getBranches(gitRepo *git.Repository) (branches []taxonomy.Branch, err error) {
	branchIter, err := gitRepo.Branches()
	if err != nil {
		return
	}
	defer func() { branchIter.Close() }()
	err = branchIter.ForEach(func(r *plumbing.Reference) error {
		if r.Type() != plumbing.HashReference {
			return fmt.Errorf("invalid reference type %s for branch %s", r.Type(), r.Name())
		}
		branches = append(branches, r)
		return nil
	})
	return
}
