package ls

import "fmt"

func Print(repos <-chan Repo) {
	for repo := range repos {
		fmt.Println(repo.path)
	}
}
