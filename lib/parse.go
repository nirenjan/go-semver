// Package semver implements the Semantic Versioning parser and manipulator
// This package conforms to the Semantic Versioning specification 2.0.0
package lib

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
	if v[0] != '.' {
		return SemVer{}, fmt.Errorf("invalid semantic version, expecting '.', got '%#v'", v[0])
	}

	semver.minor, v, err = parseInt(v[1:])
	if err != nil {
		return SemVer{}, err
	}
	if v[0] != '.' {
		return SemVer{}, fmt.Errorf("invalid semantic version, expecting '.', got '%#v'", v[0])
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
		return SemVer{}, fmt.Errorf("unexpected trailing value '%#v'", v)
	}

	return semver, nil
}

func parseInt(v string) (uint, string, error) {
	var i uint
	var val uint

	for ; v[i] >= '0' && v[i] <= '9'; i++ {
		val *= 10
		val += (uint)(v[i] - '0')

		if v[0] == '0' && i > 0 {
			return 0, v, fmt.Errorf("Leading zeroes in integer field: %#v", v)
		}
	}

	return val, v[i:], nil
}

func parsePrerelease(v string) ([]string, string, error) {
	return []string{}, v, nil
}

func parseBuild(v string) ([]string, string, error) {
	return []string{}, v, nil
}
