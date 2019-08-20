// Package semver implements the Semantic Versioning parser and manipulator
// This package conforms to the Semantic Versioning specification 2.0.0
package lib // import "nirenjan.org/semver/lib"

import (
	"fmt"
)

// SemVer holds the parsed semantic version
type SemVer struct {
	// major holds the major version number
	major uint

	// minor holds the minor version number
	minor uint

	// patch holds the patch version number
	patch uint

	// prerelease holds the prerelease information
	prerelease []string

	// build holds the build metadata
	build []string
}

// Parse parses a string containing a Semantic Version and returns the
// parsed SemVer structure, or an error if it encountered any
func Parse(version string) (SemVer, error) {
	var semver SemVer
	var v = version
	var err error

	semver.major, v, err = parseInt(v)
	if err != nil {
		return SemVer{}, err
	}
	if len(v) == 0 {
		return SemVer{}, fmt.Errorf("invalid semantic version, expecting '.', got empty string")
	}
	if v[0] != '.' {
		return SemVer{}, fmt.Errorf("invalid semantic version, expecting '.', got '%s'", string(v[0]))
	}

	semver.minor, v, err = parseInt(v[1:])
	if err != nil {
		return SemVer{}, err
	}
	if len(v) == 0 {
		return SemVer{}, fmt.Errorf("invalid semantic version, expecting '.', got empty string")
	}
	if v[0] != '.' {
		return SemVer{}, fmt.Errorf("invalid semantic version, expecting '.', got '%s'", string(v[0]))
	}

	semver.patch, v, err = parseInt(v[1:])
	if err != nil {
		return SemVer{}, err
	}

	semver.prerelease, v, err = parsePrerelease(v)
	if err != nil {
		return SemVer{}, err
	}

	semver.build, v, err = parseBuild(v)
	if err != nil {
		return SemVer{}, err
	}

	if v != "" {
		return SemVer{}, fmt.Errorf("unexpected trailing value '%s'", v)
	}

	return semver, nil
}

func parseInt(v string) (uint, string, error) {
	var i int
	var val uint

	for ; i < len(v) && v[i] >= '0' && v[i] <= '9'; i++ {
		val *= 10
		val += (uint)(v[i] - '0')

		if v[0] == '0' && i > 0 {
			return 0, v, fmt.Errorf("leading zeroes in integer field '%s'", v)
		}
	}

	if i == 0 {
		return 0, v, fmt.Errorf("invalid integer field '%s'", v)
	}

	return val, v[i:], nil
}

func hasLeadingZeroes(field string) bool {
	possibly_numeric := true
	for _, c := range field {
		if c < '0' || c > '9' {
			possibly_numeric = false
			break
		}
	}
	return possibly_numeric && len(field) > 1 && field[0] == '0'
}

func parseFields(v string, numeric bool) ([]string, string, error) {
	var field string
	var fields []string
	var i int
	var c rune

parseLoop:
	for i, c = range v {
		switch {
		case c >= 'A' && c <= 'Z':
			fallthrough
		case c >= 'a' && c <= 'z':
			fallthrough
		case c >= '0' && c <= '9':
			fallthrough
		case c == '-':
			field += string(c)
		case c == '.':
			// This is the separator field
			if field == "" {
				return []string{}, v, fmt.Errorf("unexpected '.' in dot-separated field")
			}
			if numeric && hasLeadingZeroes(field) {
				return []string{}, v, fmt.Errorf("leading zeroes in numeric field '%s' of prerelease", field)
			}
			fields = append(fields, field)
			field = "" // reset the field for the next round
		case c == '+' && numeric:
			// Accept the +, but break out of the loop and return the
			// computed fields
			i--
			break parseLoop
		default:
			return []string{}, v, fmt.Errorf("unexpected character '%s' in dot-separated field", string(c))
		}
	}

	if field != "" {
		if numeric && hasLeadingZeroes(field) {
			return []string{}, v, fmt.Errorf("leading zeroes in numeric field '%s' of prerelease", field)
		}

		fields = append(fields, field)
	} else if len(fields) > 0 {
		return []string{}, v, fmt.Errorf("invalid dot-separated field")
	}

	if len(fields) == 0 {
		return []string{}, v, fmt.Errorf("invalid dot-separated field '%s'", v)
	}

	return fields, v[i+1:], nil
}

func parsePrerelease(v string) ([]string, string, error) {
	if len(v) > 0 && v[0] == '-' {
		return parseFields(v[1:], true)
	}

	return []string{}, v, nil
}

func parseBuild(v string) ([]string, string, error) {
	if len(v) > 0 && v[0] == '+' {
		return parseFields(v[1:], false)
	}

	return []string{}, v, nil
}
