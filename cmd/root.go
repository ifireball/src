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
	"fmt"
	"os"
	"path"
	"time"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "src",
	Short: "A brief description of your application",
	Long: `"src" is a tool for managing source code directories with sources
that belong to many different projects.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/src/src.toml)")

	defaultPruneThreshold := time.Hour * 24 * 30 * 6
	rootCmd.PersistentFlags().Duration(
		"prune-threshold",
		defaultPruneThreshold,
		"A time duration to use as a pruning threshold. Repos that we last\n" +
		"committed to more then the threshold time ago are considered to be\n" +
		"'old' and may be pruned",
	)
	viper.SetDefault("prune-threshold", defaultPruneThreshold)
	viper.BindPFlag("prune-threshold", rootCmd.PersistentFlags().Lookup("prune-threshold"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile == "" {
		cfgFile = path.Join(xdg.ConfigHome, "src", "src.toml")
	}

	viper.SetConfigFile(cfgFile)

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
