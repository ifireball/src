package clone

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("getCloneLocation", func() {
	DescribeTable("Given a repo URL, returns the path to clone the repo into",
		func(repoUrl, expPath string) {
			Expect(getCloneLocation(repoUrl)).To(Equal(expPath))
		},
		EntryDescription("%[1]s"),
		Entry(nil, "https://github.com/cli/cli.git", "github.com/cli/cli"),
		Entry(nil, "git@github.com:cli/cli.git", "github.com/cli/cli"),
	)
})
