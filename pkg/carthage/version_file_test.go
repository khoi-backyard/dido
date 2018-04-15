package carthage

import (
	"testing"
)

var frameworktests = []struct {
	path             string
	reponame         string
	watchOSHashCount int
}{
	{"../../testdata/.Alamofire.version", "Alamofire", 1},
	{"../../testdata/.RxSwift.version", "RxSwift", 3},
	{"../../testdata/.swift-statsd-client.version", "swift-statsd-client", 0},
}

func TestNewVersionFile(t *testing.T) {
	for _, c := range frameworktests {
		f, err := NewVersionFile(c.path)
		if err != nil {
			t.Errorf("fail to parse %s", c.path)
		}
		if f.RepoName != c.reponame {
			t.Errorf("invalid repo name, expected %s got %s", c.reponame, f.RepoName)
		}
	}
}

func TestVersionFile_Hashes(t *testing.T) {
	for _, c := range frameworktests {
		f, err := NewVersionFile(c.path)
		if err != nil {
			t.Errorf("fail to parse %s", c.path)
		}
		if count := len(f.Hashes(PLATFORM_watchOS)); count != c.watchOSHashCount {
			t.Errorf("invalid framework hashes count, expected %d %d", c.watchOSHashCount, count)
		}
	}
}
