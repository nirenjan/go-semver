package lib

import "testing"

// compareSemvers compares two SemVer structs and returns true if they are
// both equal
func compareSemvers(s1, s2 SemVer) bool {
	if s1.major != s2.major {
		return false
	}

	if s1.minor != s2.minor {
		return false
	}

	if s1.patch != s2.patch {
		return false
	}

	if len(s1.prerelease) != len(s2.prerelease) {
		return false
	}

	for i := range s1.prerelease {
		if s1.prerelease[i] != s2.prerelease[i] {
			return false
		}
	}

	if len(s1.build) != len(s2.build) {
		return false
	}

	for i := range s1.build {
		if s1.build[i] != s2.build[i] {
			return false
		}
	}

	return true
}

func TestParsing(t *testing.T) {
	tables := []struct {
		input  string
		parsed SemVer
		err    string
	}{
		{"0.0.0", SemVer{0, 0, 0, []string{}, []string{}}, ""},
		{"0.1.2", SemVer{0, 1, 2, []string{}, []string{}}, ""},
		{"0.01.2", SemVer{}, "leading zeroes in integer field '01.2'"},
		{"0", SemVer{}, "invalid semantic version, expecting '.', got empty string"},
		{"0.", SemVer{}, "invalid integer field ''"},
		{"0+", SemVer{}, "invalid semantic version, expecting '.', got '+'"},
		{"0.0", SemVer{}, "invalid semantic version, expecting '.', got empty string"},
		{"0.0.", SemVer{}, "invalid integer field ''"},
		{"0.0+", SemVer{}, "invalid semantic version, expecting '.', got '+'"},
		{"0.0.0", SemVer{}, ""},
		{"0.0.0.", SemVer{}, "unexpected trailing value '.'"},
		{"0.0.0=", SemVer{}, "unexpected trailing value '='"},
		{"foo", SemVer{}, "invalid integer field 'foo'"},
		{".", SemVer{}, "invalid integer field '.'"},
	}

	for _, table := range tables {
		parsed, err := Parse(table.input)
		if !compareSemvers(parsed, table.parsed) {
			t.Errorf("Parsing '%#v' parse mismatch: expected '%#v', got '%#v'",
				table.input, table.parsed, parsed)
		}
		if err == nil && table.err != "" {
			t.Errorf("Parsing '%#v' error mismatch: expected '%#v', got nil",
				table.input, table.err)
		}
		if err != nil && err.Error() != table.err {
			t.Errorf("Parsing '%#v' error mismatch: expected '%#v', got '%#v'",
				table.input, table.err, err.Error())
		}
	}
}
