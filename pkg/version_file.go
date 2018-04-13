package pkg

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"path/filepath"

	"errors"
)

const VersionFileExtension = ".version"

var ErrInvalidFileName = errors.New("invalid filename format")

type VersionFile struct {
	RepoName string
	Version  *VersionFileContent
}

type VersionHash struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

type VersionFileContent struct {
	Commitish string        `json:"commitish"`
	Mac       []VersionHash `json:"Mac"`
	WatchOS   []VersionHash `json:"watchOS"`
	TvOS      []VersionHash `json:"tvOS"`
	IOS       []VersionHash `json:"iOS"`
}

// NewVersionFile parse the .version file and return the model
func NewVersionFile(path string) (*VersionFile, error) {
	_, fileName := filepath.Split(path)
	if fileName[0] != '.' || !strings.HasSuffix(fileName, VersionFileExtension) {
		return nil, ErrInvalidFileName
	}

	b, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var version VersionFileContent
	err = json.Unmarshal(b, &version)
	if err != nil {
		return nil, err
	}

	return &VersionFile{
		fileName[1 : len(fileName)-len(filepath.Ext(fileName))],
		&version,
	}, nil
}

func (f *VersionFile) Hashes(platform string) []VersionHash {
	switch strings.ToLower(platform) {
	case PLATFORM_iOS:
		return f.Version.IOS
	case PLATFORM_macOS:
		return f.Version.Mac
	case PLATFORM_tvOS:
		return f.Version.TvOS
	case PLATFORM_watchOS:
		return f.Version.WatchOS
	default:
		return nil
	}
}
