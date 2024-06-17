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

	"github.com/ifireball/src/lib/clone"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone URL",
	Short: "Clone a given repo to the right location under ~/src",
	Long: `Clone a repo given by its clone url to the right location 
under ~/src. Leading directories are created as needed.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		src := path.Join(home, "src")
		srcFs := afero.NewOsFs()
		cobra.CheckErr(clone.Repo(srcFs, src, args[0]))
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
}
