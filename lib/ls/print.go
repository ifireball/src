package ls

import (
	"fmt"
	"slices"
	"strings"
)

type tableRow struct{name, since, remote string}

func Print(repos <-chan Repo, showRemote bool) {
	var table []tableRow
	nameWidth := 0
	for repo := range repos {
		name := fmt.Sprintf("%s/%s/%s", repo.Host, repo.Org, repo.Name)
		nameWidth = max(nameWidth, len(name))
		table = append(table, tableRow{
			name,
			repo.LastCommitTime.Format("Mon, Jan _2 2006"),
			repo.MainRemoteURL,
		})
	}
	slices.SortFunc(table, func(a, b tableRow) int {
		return strings.Compare(a.name, b.name)
	})
	for _, row := range table {
		fmt.Printf("%-[1]*[2]s @ %[3]s\n", nameWidth, row.name, row.since)
		if showRemote {
			fmt.Printf("  url: %s\n", row.remote)
		}
	}
}
