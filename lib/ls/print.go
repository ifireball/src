package ls

import (
	"fmt"
	"slices"
	"strings"
)

type tableRow struct{name, since string}

func Print(repos <-chan Repo) {
	var table []tableRow
	nameWidth := 0
	for repo := range repos {
		name := fmt.Sprintf("%s/%s/%s", repo.Host, repo.Org, repo.Name)
		nameWidth = max(nameWidth, len(name))
		table = append(table, struct{name string; since string}{
			name,
			repo.LastCommitTime.Format("Mon, Jan _2 2006"),
		})
	}
	slices.SortFunc(table, func(a, b tableRow) int {
		return strings.Compare(a.name, b.name)
	})
	for _, row := range table {
		fmt.Printf("%-[1]*[2]s @ %[3]s\n", nameWidth, row.name, row.since)
	}
}
