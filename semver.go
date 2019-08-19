// Package semver implements the Semantic Versioning parser and manipulator
// This package conforms to the Semantic Versioning specification 2.0.0
package semver

import (
	"nirenjan.org/semver/lib"
)

// SemVer holds the parsed semantic version
type SemVer struct {
	semver lib.SemVer
}

// Parse parses a string containing a Semantic Version and returns the
// parsed SemVer structure, or an error if it encountered any
func Parse(version string) (SemVer, error) {
	semver, err := lib.Parse(version)

	if err != nil {
		return SemVer{}, err
	}
	return SemVer{semver}, nil
}
