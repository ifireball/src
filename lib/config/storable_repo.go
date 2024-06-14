package config

import (
	"github.com/go-git/go-git/v5/config"

	"github.com/ifireball/src/lib/taxonomy"
)

// Interface of repos we can store
type StorableRepo interface {
	ShortPath() string
	Config() *config.Config
	Branches() []taxonomy.Branch
	Head() string
}
