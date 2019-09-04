package semver

import "fmt"

func formatDotSeparated(s []string) string {
	var out string

	for _, v := range s {
		out += fmt.Sprintf(".%s", v)
	}

	// Drop the leading .
	return out[1:]
}

func (s SemVer) String() string {
	var out string

	out = fmt.Sprintf("%d.%d.%d", s.major, s.minor, s.patch)

	if len(s.prerelease) > 0 {
		out += "-" + formatDotSeparated(s.prerelease)
	}

	if len(s.build) > 0 {
		out += "+" + formatDotSeparated(s.build)
	}

	return out
}
