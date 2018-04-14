package pkg

import (
	"errors"
	"os/user"
	"path/filepath"
)

var ErrCannotExpandPath = errors.New("can't expand path")

// Expand expands the path prefixed by `~`
func Expand(path string) (string, error) {
	if len(path) == 0 || path[0] != '~' {
		return path, nil
	}

	if len(path) > 1 && path[1] != '/' && path[1] != '\\' {
		return "", ErrCannotExpandPath
	}

	usr, err := user.Current()

	if err != nil {
		return "", err
	}

	return filepath.Join(usr.HomeDir, path[1:]), nil
}
