package clone

import (
	"fmt"
	"net/url"
	"path"
	"strings"
)

// Get the location to clone a repo into
func getCloneLocation(repoUrl string) (string, error) {
	url, err := url.Parse(repoUrl)
	if err != nil {
		return getSSHURLCloneLocation(repoUrl)
	}
	repoDir := strings.TrimSuffix(url.Path, ".git")
	return url.Host + repoDir, nil
}

func getSSHURLCloneLocation(repoUrl string) (string, error) {
	urlParts := strings.SplitN(repoUrl, ":", 3)
	if len(urlParts) != 2 {
		return "", fmt.Errorf("invalid Git SSH URL: %s", repoUrl)
	}
	userHostParts := strings.SplitN(urlParts[0], "@", 3)
	if len(userHostParts) > 2 {
		return "", fmt.Errorf("invalid Git SSH URL: %s", repoUrl)
	}
	repoHost := userHostParts[len(userHostParts)-1]
	repoDir := strings.TrimSuffix(urlParts[1], ".git")
	if path.IsAbs(repoDir) {
		return repoHost + repoDir, nil
	} else {
		return repoHost + "/" + repoDir, nil
	}
}
