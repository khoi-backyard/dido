package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"path/filepath"

	"github.com/pkg/errors"
)

const VersionFileExtension = ".version"

var ErrInvalidFileName = errors.New("invalid filename format")

type Framework struct {
	RepoName string
	Version  *Version
}

type Version struct {
	Commitish string `json:"commitish"`
	Mac       []struct {
		Name string `json:"name"`
		Hash string `json:"hash"`
	} `json:"Mac"`
	WatchOS []struct {
		Name string `json:"name"`
		Hash string `json:"hash"`
	} `json:"watchOS"`
	TvOS []struct {
		Name string `json:"name"`
		Hash string `json:"hash"`
	} `json:"tvOS"`
	IOS []struct {
		Name string `json:"name"`
		Hash string `json:"hash"`
	} `json:"iOS"`
}

// NewFramework parse the .version file and return the model
func NewFramework(path string) (*Framework, error) {
	_, fileName := filepath.Split(path)
	if fileName[0] != '.' || !strings.HasSuffix(fileName, VersionFileExtension) {
		return nil, ErrInvalidFileName
	}

	b, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	var version Version
	err = json.Unmarshal(b, &version)
	if err != nil {
		return nil, err
	}

	return &Framework{
		fileName[1 : len(fileName)-len(filepath.Ext(fileName))],
		&version,
	}, nil
}

func (f *Framework) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s \n", f.RepoName, f.Version.Commitish))
	if len(f.Version.Mac) > 0 {
		buffer.WriteString(fmt.Sprintf("   Mac: \n"))
		for _, f := range f.Version.Mac {
			buffer.WriteString(fmt.Sprintf("      %s %s \n", f.Name, f.Hash))
		}
	}
	if len(f.Version.IOS) > 0 {
		buffer.WriteString(fmt.Sprintf("   iOS: \n"))
		for _, f := range f.Version.IOS {
			buffer.WriteString(fmt.Sprintf("      %s %s \n", f.Name, f.Hash))
		}
	}
	if len(f.Version.TvOS) > 0 {
		buffer.WriteString(fmt.Sprintf("   tvOS: \n"))
		for _, f := range f.Version.TvOS {
			buffer.WriteString(fmt.Sprintf("      %s %s \n", f.Name, f.Hash))
		}
	}
	if len(f.Version.WatchOS) > 0 {
		buffer.WriteString(fmt.Sprintf("   watchOS: \n"))
		for _, f := range f.Version.WatchOS {
			buffer.WriteString(fmt.Sprintf("      %s %s \n", f.Name, f.Hash))
		}
	}
	return buffer.String()
}
