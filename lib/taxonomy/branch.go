package taxonomy

import "github.com/go-git/go-git/v5/plumbing"

type Branch interface {
	Name() plumbing.ReferenceName
	Hash() plumbing.Hash
}
