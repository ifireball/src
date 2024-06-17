package clone

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/afero"
)

func Repo(srcFS afero.Fs, srcDir string, repoURL string) error {
	location, err := getCloneLocation(repoURL)
	if err != nil {
		return err
	}
	fullPath := path.Join(srcDir, location)
	if _, err := srcFS.Stat(fullPath); !os.IsNotExist(err) {
		if err == nil {
			return fmt.Errorf("target directory %s exists, skipping cloning", location)
		} else {
			return err
		}
	}
	fmt.Printf("Cloning repository into: %s\n", location)
	if err := srcFS.MkdirAll(path.Dir(fullPath), fs.ModeDir|0755); err != nil {
		return err
	}
	_, err = git.PlainClone(fullPath, false, &git.CloneOptions{URL: repoURL})
	return err
}
