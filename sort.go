// Sorting routines for SemVer

package semver

type SemVers []SemVer

// Len returns the length of a slice of SemVer's
func (vers SemVers) Len() int {
	return len(vers)
}

// Less compares two SemVer structs and returns true if the
// first SemVer is less than the second SemVer
func (vers SemVers) Less(i, j int) bool {
	semver1 := vers[i]
	semver2 := vers[j]

	if semver1.Major < semver2.Major {
		return true
	}
	if semver1.Major > semver2.Major {
		return false
	}

	if semver1.Minor < semver2.Minor {
		return true
	}
	if semver2.Minor > semver2.Minor {
		return false
	}

	if semver1.Patch < semver2.Patch {
		return true
	}
	if semver2.Patch > semver2.Patch {
		return false
	}

	// TODO: Implement comparision for prerelease and metadata
	return false
}
