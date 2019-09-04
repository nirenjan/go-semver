package semver

import "testing"

// TestCompare compares two Semantic versions and matches the
// expected result against the desired result
func TestCompare(t *testing.T) {
	tables := []struct {
		v1  string
		v2  string
		exp int
	}{
		{"0.0.0", "0.0.0", 0},
		{"0.0.0", "0.0.1", -1},
		{"0.0.1", "0.0.0", 1},
		{"0.1.1", "0.1.0", 1},
		{"1.0.0", "0.1.0", 1},
		{"1.0.0", "1.1.0", -1},

		{"1.0.0", "1.0.0-alpha", 1},
		{"1.0.0", "1.0.0+alpha", 0},

		{"1.0.0-alpha", "1.0.0+alpha", -1},

		{"1.0.0-alpha", "1.0.0-0", 1},
		{"1.0.0-alpha", "1.0.0-0.alpha", 1},
		{"1.0.0-0.alpha", "1.0.0-0.alpha", 0},
		{"1.0.0-0.alpha", "1.0.0-0.alpha.1", -1},

		{"1.0.0", "2.0.0", -1},
		{"2.0.0", "2.1.0", -1},
		{"2.1.0", "2.1.1", -1},

		{"1.0.0-alpha", "1.0.0-alpha.1", -1},
		{"1.0.0-alpha.1", "1.0.0-alpha.beta", -1},
		{"1.0.0-alpha.beta", "1.0.0-beta", -1},
		{"1.0.0-beta", "1.0.0-beta.2", -1},
		{"1.0.0-beta.2", "1.0.0-beta.11", -1},
		{"1.0.0-beta.11", "1.0.0-rc.1", -1},
		{"1.0.0-rc.1", "1.0.0", -1},

		{"1.0.0-b", "1.0.0-a", 1},

		{"1.0.0-alpha", "1.0.0-alpha+001", 0},
		{"1.0.0", "1.0.0+20130313144700", 0},
		{"1.0.0-beta", "1.0.0-beta+exp.sha.5114f85", 0},
	}

	for _, tc := range tables {
		s1, _ := Parse(tc.v1)
		s2, _ := Parse(tc.v2)

		if got := s1.Compare(s2); got != tc.exp {
			t.Errorf("Comparing '%s' to '%s', expected '%d', got '%d'",
				s1, s2, tc.exp, got)
		}
	}
}
