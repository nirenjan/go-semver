package semver

// SemVer is the default structure which contains the fields Major, Minor,
// Patch, and optional Prerelease and Metadata
type SemVer struct {
	Major      uint     // Major version number
	Minor      uint     // Minor version number
	Patch      uint     // Patch version number
	Prerelease []string // Prerelease information
	Metadata   []string // Build metadata
}

func emptySemVer() SemVer {
	return SemVer{0, 0, 0, make([]string, 0), make([]string, 0)}
}
