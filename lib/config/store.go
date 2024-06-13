package config

import (
	"slices"
	"strings"

	"github.com/ifireball/src/lib/ls"
	"github.com/spf13/viper"
)

type repoConfig struct {
	Path, RemoteURL string
}

func Store(repos <-chan ls.Repo, v *viper.Viper) {
	var repoConfigs []repoConfig
	for repo := range repos {
		repoConfigs = append(repoConfigs, repoConfig{
			Path: repo.ShortPath(),
			RemoteURL: repo.MainRemoteURL,
		})
	}
	slices.SortFunc(repoConfigs, func(a, b repoConfig) int {
		return strings.Compare(a.Path, b.Path)
	})
	v.Set("repositories", repoConfigs)
}
