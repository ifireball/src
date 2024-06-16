package config

import "github.com/spf13/viper"

func Load(v *viper.Viper) (<-chan StorableRepo, error) {
	var repoConfigs []repoConfig
	out  := make(chan StorableRepo)
	err := v.UnmarshalKey("repositories", repoConfigs)
	go func() {
		defer func() { close(out) }()
		for _, repoConfig := range repoConfigs {
			out <- &repoConfig
		}
	}()
	return out, err
}
