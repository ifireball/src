package ls

import "fmt"

func Print(repos <-chan Repo) {
	for repo := range repos {
		fmt.Printf("%s/%s/%s @ %s\n", repo.Host, repo.Org, repo.Name, repo.LastCommitTime)
	}
}
