package pkg

import "testing"

var tests = []struct {
	in       string
	origin   string
	location string
	repoName string
	version  string
}{
	{`github "Alamofire/Alamofire" "4.7.0"`, "github", "Alamofire/Alamofire", "Alamofire", "4.7.0"},
	{`github "Cocoanetics/DTCoreText" "1.7.13"`, "github", "Cocoanetics/DTCoreText", "DTCoreText", "1.7.13"},
	{`github "Instagram/IGListKit" "3.1.1"`, "github", "Instagram/IGListKit", "IGListKit", "3.1.1"},
	{`github "rs/SDWebImage" "85665a5af97c728bb64c8c07c4021488b237b164"`, "github", "rs/SDWebImage", "SDWebImage", "85665a5af97c728bb64c8c07c4021488b237b164"},
	{`git "https://enterprise.local/desktop/git-error-translations2.git" "d681925d54076aa921b6ab7a0691cc99f6098fe7"`, "git", "https://enterprise.local/desktop/git-error-translations2.git", "git-error-translations2", "d681925d54076aa921b6ab7a0691cc99f6098fe7"},
}

func TestNewResolution(t *testing.T) {
	for _, tt := range tests {
		r, err := NewResolution(tt.in)
		if err != nil {
			t.Errorf("Failed to parse line %s", tt.in)
		}

		if r.Origin != tt.origin {
			t.Errorf("Wrong origin, expecting %s got %s", tt.origin, r.Origin)
		}

		if r.Location != tt.location {
			t.Errorf("Wrong location, expecting %s got %s", tt.location, r.Location)
		}

		if r.Version != tt.version {
			t.Errorf("Wrong version, expecting %s got %s", tt.version, r.Version)
		}

		if r.RepoName() != tt.repoName {
			t.Errorf("Wrong reponame, expecting %s got %s", tt.repoName, r.RepoName())
		}
	}
}
