package ls

import (
	"strings"
)

type RepoPathData interface {
	ShortPath() string
}

type repoPathDataImpl struct {
	host, org, name string
}

func getRepoPathData(srcPath, repo string) *repoPathDataImpl {
	repo = strings.TrimPrefix(repo, srcPath)
	repo = strings.TrimPrefix(repo, "/")
	repo = strings.TrimSuffix(repo, "/")
	firstI := strings.Index(repo, "/")
	lastI := strings.LastIndex(repo, "/")

	var host, org, name string
	if firstI >= 0 {
		host = repo[:firstI]
	}
	if firstI < lastI {
		org = repo[firstI+1:lastI]
	}
	name = repo[lastI+1:]
	return &repoPathDataImpl{host: host, org: org, name: name}
}

func (rpd *repoPathDataImpl) ShortPath() string {
	if rpd.host != "" {
		if rpd.org != "" {
			return strings.Join([]string{rpd.host, rpd.org, rpd.name}, "/")
		}
		return strings.Join([]string{rpd.host, rpd.name}, "/")
	}
	return rpd.name
}
