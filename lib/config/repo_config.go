package config

import (
	"github.com/go-git/go-git/v5/config"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/ifireball/src/lib/taxonomy"
)

// A branch representation that marshals nicely
type branchConfig struct {
	HashStr string `mapstructure:"hash"`
	NameStr string `mapstructure:"name"`
}

// A repo representation that marshals nicely
type repoConfig struct {
	PathData string            `mapstructure:"path"`
	ConfigData string          `mapstructure:"config"`
	BranchData []*branchConfig `mapstructure:"branches"`
	HeadRef string             `mapstructure:"head"`
}

func repoConfigFromStorable(repo StorableRepo) (*repoConfig, error) {
	cfg, err := repo.Config().Marshal()
	if err != nil {
		return nil, err
	}
	config := repoConfig{
		PathData: repo.ShortPath(),
		ConfigData: string(cfg),
		BranchData: storeBranches(repo.Branches()),
		HeadRef: repo.Head(),
	}
	return &config, nil
}

func storeBranches(branches []taxonomy.Branch) (branchConfigs []*branchConfig) {
	for _, branch := range branches {
		branchConfigs = append(branchConfigs, &branchConfig{
			HashStr: branch.Hash().String(),
			NameStr: branch.Name().String(),
		})
	}
	return
}

func (r *repoConfig) ShortPath() string {
	return r.PathData
}

func (r* repoConfig) Config() *config.Config {
	cfg := config.NewConfig()
	cfg.Unmarshal([]byte(r.ConfigData))
	return cfg
}

func (r *repoConfig) Branches() []taxonomy.Branch {
	out := make([]taxonomy.Branch, len(r.BranchData))
	for i, branch := range r.BranchData {
		out[i] = branch
	}
	return out
}

func (r *repoConfig) Head() string { return r.HeadRef }

func (b *branchConfig) Hash() plumbing.Hash {
	return plumbing.NewHash(b.HashStr)
}

func (b *branchConfig) Name() plumbing.ReferenceName {
	return plumbing.ReferenceName(b.NameStr)
}
