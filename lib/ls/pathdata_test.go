package ls

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("getRepoPathData", func() {
	srcPath := "/path/to/src"
	DescribeTable("Extracts repo host, org and name from path",
		func(repo, expHost, expOrg, expName string) {
			expected := RepoPathData{Host: expHost, Org: expOrg, Name: expName}
			Expect(getRepoPathData(srcPath, repo)).To(Equal(expected))
		},
		EntryDescription("%[1]s"),
		Entry(nil, "/path/to/src/host/org/repo", "host", "org", "repo"),
		Entry(nil, "/path/to/src/host/org/repo/", "host", "org", "repo"),
		Entry(nil, "/path/to/src/host/org/sub/repo", "host", "org/sub", "repo"),
		Entry(nil, "/path/to/src/host/org/sub/repo/", "host", "org/sub", "repo"),
		Entry(nil, "/path/to/src/host/repo", "host", "", "repo"),
		Entry(nil, "/path/to/src/host/repo/", "host", "", "repo"),
		Entry(nil, "/path/to/src/repo", "", "", "repo"),
		Entry(nil, "/path/to/src/repo/", "", "", "repo"),
	)
})
