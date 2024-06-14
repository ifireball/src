package config

import (
	"slices"
	"strings"

	"github.com/spf13/viper"
)

func Store[T StorableRepo](repos <-chan T, v *viper.Viper) error {
	var repoConfigs []repoConfig
	for repo := range repos {
		cfg, err := repoConfigFromStorable(repo)
		if err != nil {
			return err
		}
		repoConfigs = append(repoConfigs, *cfg)
	}
	slices.SortFunc(repoConfigs, func(a, b repoConfig) int {
		return strings.Compare(a.Path, b.Path)
	})
	v.Set("repositories", repoConfigs)
	return nil
}
