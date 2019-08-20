package lib

import "testing"

// TestStringify parses a given string, and checks if the
// String output matches the input
func TestStringify(t *testing.T) {
	cases := []string{
		"0.0.0",
		"0.1.2",
		"13.37.1337",
		"1.2.3",
		"1.2.3-a",
		"1.2.3-a.10",
		"1.2.3-a.10.b",
		"1.2.3-a.0.b",
		"1.2.3-beta1",
		"1.2.3-0beta1",
		"1.2.3+g",
		"1.2.3+gabcdef0",
		"1.2.3+gabcdef0.0012",
		"1.2.3+UPPERCASE",
		"1.2.3+b1",
		"1.2.3-alpha.1+build.0820.gabcdef0",
	}

	for _, v := range cases {
		s, _ := Parse(v)

		if v != s.String() {
			t.Errorf("Mismatch in String, expected '%s', got '%s'", v, s.String())
		}
	}
}
