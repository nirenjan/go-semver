package semver

type SemVer struct {
	major      uint
	minor      uint
	patch      uint
	prerelease []string
	metadata   []string
}

func emptySemVer() SemVer {
	return SemVer{0, 0, 0, make([]string, 0), make([]string, 0)}
}
