package pkg

import (
	"testing"
)

var frameworktests = []struct {
	path     string
	reponame string
}{
	{"../testdata/.Alamofire.version", "Alamofire"},
	{"../testdata/.RxSwift.version", "RxSwift"},
	{"../testdata/.swift-statsd-client.version", "swift-statsd-client"},
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
