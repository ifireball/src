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

	"github.com/ifireball/src/lib/config"
	"github.com/ifireball/src/lib/ls"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Store repo configuration to the configuration file",
	Long: `Store the list of repos and any other configuration settings you've
specified into the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		src := path.Join(home, "src")
		srcFs := os.DirFS("/")
		repos, err := ls.Repos(srcFs, src)
		cobra.CheckErr(err)
		config.Store(repos, viper.GetViper())
		configFile := viper.ConfigFileUsed()
		cobra.CheckErr(viper.WriteConfigAs(configFile))
	},
}

func init() {
	configCmd.AddCommand(storeCmd)
}
