package pkg

import (
	"os/user"
	"path/filepath"
	"testing"
)

func TestExpand(t *testing.T) {
	u, err := user.Current()
	if err != nil {
		t.Fatalf("%s", err)
	}
	tests := []struct {
		in  string
		out string
	}{
		{"/tmp/", "/tmp/"},
		{"/tmp/dido", "/tmp/dido"},
		{"~/Library/Cache", filepath.Join(u.HomeDir, "Library/Cache")},
		{"~asdf/Library/Cache", ""},
	}

	for _, c := range tests {
		result, _ := Expand(c.in)
		if result != c.out {
			t.Errorf("Expect %s - Got %s", c.out, result)
		}
	}
}
