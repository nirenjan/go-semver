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

		{"13.37.1337", SemVer{13, 37, 1337, []string{}, []string{}}, ""},
		{"1.2.3", SemVer{1, 2, 3, []string{}, []string{}}, ""},
		{"1.2.3-a", SemVer{1, 2, 3, []string{"a"}, []string{}}, ""},
		{"1.2.3-a.10", SemVer{1, 2, 3, []string{"a", "10"}, []string{}}, ""},
		{"1.2.3-a.10.b", SemVer{1, 2, 3, []string{"a", "10", "b"}, []string{}}, ""},
		{"1.2.3-a.0.b", SemVer{1, 2, 3, []string{"a", "0", "b"}, []string{}}, ""},
		{"1.2.3-beta1", SemVer{1, 2, 3, []string{"beta1"}, []string{}}, ""},
		{"1.2.3-0beta1", SemVer{1, 2, 3, []string{"0beta1"}, []string{}}, ""},
		{"1.2.3-a..b", SemVer{}, "unexpected '.' in dot-separated field"},
		{"1.2.3-a.10.b.", SemVer{}, "invalid dot-separated field"},
		{"1.2.3-a.010", SemVer{}, "leading zeroes in numeric field '010' of prerelease"},
		{"1.2.3-a.00.1", SemVer{}, "leading zeroes in numeric field '00' of prerelease"},
		{"1.2.3-", SemVer{}, "invalid dot-separated field ''"},
		{"1.2.3-.1", SemVer{}, "unexpected '.' in dot-separated field"},

		{"1.2.3+g", SemVer{1, 2, 3, []string{}, []string{"g"}}, ""},
		{"1.2.3+gabcdef0", SemVer{1, 2, 3, []string{}, []string{"gabcdef0"}}, ""},
		{"1.2.3+gabcdef0.0012", SemVer{1, 2, 3, []string{}, []string{"gabcdef0", "0012"}}, ""},
		{"1.2.3+UPPERCASE", SemVer{1, 2, 3, []string{}, []string{"UPPERCASE"}}, ""},
		{"1.2.3+b1", SemVer{1, 2, 3, []string{}, []string{"b1"}}, ""},

		{"1.2.3-alpha.1+build.0820.gabcdef0", SemVer{1, 2, 3, []string{"alpha", "1"}, []string{"build", "0820", "gabcdef0"}}, ""},

		{"1.2.3+alpha.1%", SemVer{}, "unexpected character '%' in dot-separated field"},

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
