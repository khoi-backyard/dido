package carthage

import (
	"io/ioutil"
	"path"
	"strings"
)

func GetVersionFiles(buildFolderPath string) ([]*VersionFile, error) {
	files, err := ioutil.ReadDir(buildFolderPath)

	if err != nil {
		return nil, err
	}

	var versionFiles []*VersionFile

	for _, f := range files {
		if !strings.HasSuffix(f.Name(), VersionFileExtension) {
			continue
		}

		versionFilePath := path.Join(buildFolderPath, f.Name())
		versionFile, err := NewVersionFile(versionFilePath)

		if err != nil {
			return nil, err
		}

		versionFiles = append(versionFiles, versionFile)
	}

	return versionFiles, nil
}
