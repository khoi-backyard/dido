package carthage

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

const FrameworkExt = ".framework"
const DsymExt = ".dSYM"

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

// FrameworkExist returns the path, if the framework exists on disk
func FrameworkExist(buildFolderPath, name, platform string) (string, error) {
	frameworkPath := path.Join(buildFolderPath, platform, fmt.Sprintf("%s%s", name, FrameworkExt))
	if _, err := os.Stat(frameworkPath); os.IsNotExist(err) {
		return "", err
	}
	return frameworkPath, nil
}

func DsymExist(buildFolderPath, name, platform string) (string, error) {
	dsymPath := path.Join(buildFolderPath, platform, fmt.Sprintf("%s%s%s", name, FrameworkExt, DsymExt))
	if _, err := os.Stat(dsymPath); os.IsNotExist(err) {
		return "", err
	}
	return dsymPath, nil
}
