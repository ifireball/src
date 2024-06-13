package out

import (
	"fmt"
	"slices"
	"strings"
	"time"
)

type tableRow struct{name, since, remote string}

type PrintableRepo interface {
	ShortPath() string
	LastCommitTime() time.Time
	MainRemoteURL() string
}

func Print[T PrintableRepo](repos <-chan T, showRemote bool) {
	var table []tableRow
	nameWidth := 0
	for repo := range repos {
		name := repo.ShortPath()
		nameWidth = max(nameWidth, len(name))
		table = append(table, tableRow{
			name,
			repo.LastCommitTime().Format("Mon, Jan _2 2006"),
			repo.MainRemoteURL(),
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
