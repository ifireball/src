package ls

import (
	"strings"
)


type RepoPathData struct {
	Host, Org, Name string
}

func getRepoPathData(srcPath, repo string) RepoPathData {
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
	return RepoPathData{Host: host, Org: org, Name: name}
}

func (rpd *RepoPathData) ShortPath() string {
	if rpd.Host != "" {
		if rpd.Org != "" {
			return strings.Join([]string{rpd.Host, rpd.Org, rpd.Name}, "/")
		}
		return strings.Join([]string{rpd.Host, rpd.Name}, "/")
	}
	return rpd.Name
}
