package config

import (
	"slices"
	"strings"

	"github.com/spf13/viper"
)

type StorableRepo interface {
	ShortPath() string
	MainRemoteURL() string
}

type repoConfig struct {
	Path, RemoteURL string
}

func Store[T StorableRepo](repos <-chan T, v *viper.Viper) {
	var repoConfigs []repoConfig
	for repo := range repos {
		repoConfigs = append(repoConfigs, repoConfig{
			Path: repo.ShortPath(),
			RemoteURL: repo.MainRemoteURL(),
		})
	}
	slices.SortFunc(repoConfigs, func(a, b repoConfig) int {
		return strings.Compare(a.Path, b.Path)
	})
	v.Set("repositories", repoConfigs)
}
