package config

import "github.com/ifireball/src/lib/taxonomy"

// A branch representation that marshals nicely
type branchConfig struct {
	Hash, Name string
}

// A repo representation that marshals nicely
type repoConfig struct {
	Path, Config string
	Branches []*branchConfig
	Head string
}

func repoConfigFromStorable(repo StorableRepo) (*repoConfig, error) {
	cfg, err := repo.Config().Marshal()
	if err != nil {
		return nil, err
	}
	config := repoConfig{
		Path: repo.ShortPath(),
		Config: string(cfg),
		Branches: storeBranches(repo.Branches()),
		Head: repo.Head(),
	}
	return &config, nil
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
