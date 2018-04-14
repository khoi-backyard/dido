package carthage

import "testing"

var tests = []struct {
	in  string
	out bool
}{
	{"ios", true},
	{"iOs", true},
	{"tvos", true},
	{"macOS", true},
	{"watcHos", true},
	{"xcode", false},
	{"iosios", false},
}

func TestIsValidPlatform(t *testing.T) {
	for _, c := range tests {
		if output := IsValidPlatform(c.in); output != c.out {
			t.Errorf("%s Expected %t - Got %t", c.in, c.out, output)
		}
	}
}
