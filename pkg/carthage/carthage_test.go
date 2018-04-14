package carthage

import (
	"io/ioutil"
	"testing"
)

const testDataPath = "../../testdata"

func TestGetVersionFiles(t *testing.T) {
	versionFiles, err := GetVersionFiles(testDataPath)

	if err != nil {
		t.Errorf("err %s", err)
	}

	files, _ := ioutil.ReadDir(testDataPath)

	if len(files) != len(versionFiles) {
		t.Errorf("Expect %d version files - got %d", len(files), len(versionFiles))
	}
}
