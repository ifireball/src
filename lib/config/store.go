package config

import (
	"slices"
	"strings"

	"github.com/go-git/go-git/v5/config"
	"github.com/spf13/viper"
)

type StorableRepo interface {
	ShortPath() string
	Config() *config.Config
}

type repoConfig struct {
	Path string
	Config *config.Config
}

func Store[T StorableRepo](repos <-chan T, v *viper.Viper) {
	var repoConfigs []repoConfig
	for repo := range repos {
		repoConfigs = append(repoConfigs, repoConfig{
			Path: repo.ShortPath(),
			Config: repo.Config(),
		})
	}
	slices.SortFunc(repoConfigs, func(a, b repoConfig) int {
		return strings.Compare(a.Path, b.Path)
	})
	v.Set("repositories", repoConfigs)
}
