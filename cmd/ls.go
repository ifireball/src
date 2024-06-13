/*
Copyright © 2024 Barak Korren <barak.korren@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"os"
	"path"
	"time"

	"github.com/ifireball/src/lib/chu"
	"github.com/ifireball/src/lib/ls"
	"github.com/ifireball/src/lib/out"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List source code repositories",
	Long:  `List all the source code repositories you have clone into your source directory`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		src := path.Join(home, "src")
		srcFs := os.DirFS("/")
		repos, err := ls.Repos(srcFs, src)
		cobra.CheckErr(err)
		new, err := cmd.Flags().GetBool("new")
		cobra.CheckErr(err)
		old, err := cmd.Flags().GetBool("old")
		cobra.CheckErr(err)
		pruneThreshold := viper.GetDuration("prune-threshold")
		if new {
			repos = chu.Filter(repos, func(r ls.Repo) bool {
				return time.Since(r.LastCommitTime()) <= pruneThreshold
			})
		} else if old {
			repos = chu.Filter(repos, func(r ls.Repo) bool {
				return time.Since(r.LastCommitTime()) > pruneThreshold
			})
		}
		showRemote, err := cmd.Flags().GetBool("remote")
		cobra.CheckErr(err)
		out.Print(repos, showRemote)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)

	lsCmd.Flags().BoolP("old", "o", false, "Only show older repos")
	lsCmd.Flags().BoolP("new", "n", false, "Only show newer repos")
	lsCmd.MarkFlagsMutuallyExclusive("old", "new")

	lsCmd.Flags().Bool("remote", false, "Show the remote URL for repos")
}
