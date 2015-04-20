package semver

import (
	"fmt"
	"regexp"
	"strings"
)

// Prerelease parses the givem prerelease string and saves it to the
// given SemVer
func (ver *SemVer) Prerelease(prestr string) error {
	var pre = strings.Split(prestr, ".")

	var alldigits, _ = regexp.Compile("^[[:digit:]]$")
	var noleading0, _ = regexp.Compile("^(0|[1-9][[:digit:]]*)$")

	for index, comp := range pre {
		if len(comp) == 0 {
			// Disallow empty fields
			return fmt.Errorf("empty prerelease field at %v", index)
		}

		if match := alldigits.Match([]byte(comp)); match {
			if match := noleading0.Match([]byte(comp)); !match {
				// Leading zeroes are not allowed in prerelease
				return fmt.Errorf("Leading zeroes in prerelease field %q at %v", comp, index)
			}
		}
	}

	ver.prerelease = pre

	return nil
}

// Metadata converts the given dot separated build metadata string and
// saves it in the given SemVer
func (ver *SemVer) Metadata(metastr string) error {
	var metadata = strings.Split(metastr, ".")

	for index, comp := range metadata {
		if len(comp) == 0 {
			// Disallow empty fields
			return fmt.Errorf("empty metadata field at %v", index)
		}
	}

	ver.metadata = metadata

	return nil
}
