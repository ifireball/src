package config

import (
	"slices"
	"strings"

	"github.com/go-git/go-git/v5/config"
	"github.com/ifireball/src/lib/taxonomy"
	"github.com/spf13/viper"
)

type StorableRepo interface {
	ShortPath() string
	Config() *config.Config
	Branches() []taxonomy.Branch
	Head() string
}

type branchConfig struct {
	Hash, Name string
}

type repoConfig struct {
	Path, Config string
	Branches []*branchConfig
	Head string
}

func Store[T StorableRepo](repos <-chan T, v *viper.Viper) error {
	var repoConfigs []repoConfig
	for repo := range repos {
		cfg, err := repo.Config().Marshal()
		if err != nil {
			return err
		}
		repoConfigs = append(repoConfigs, repoConfig{
			Path: repo.ShortPath(),
			Config: string(cfg),
			Branches: storeBranches(repo.Branches()),
			Head: repo.Head(),
		})
	}
	slices.SortFunc(repoConfigs, func(a, b repoConfig) int {
		return strings.Compare(a.Path, b.Path)
	})
	v.Set("repositories", repoConfigs)
	return nil
}

func storeBranches(branches []taxonomy.Branch) (branchConfigs []*branchConfig) {
	for _, branch := range branches {
		branchConfigs = append(branchConfigs, &branchConfig{
			Hash: branch.Hash().String(),
			Name: branch.Name().String(),
		})
	}
	return
}
