package semver

import (
	"fmt"
	"strings"
)

// String converts the semver argument into a string of the form
// <major>.<minor>.<patch>[-<pre.release>][+<meta.data>]
func String(ver SemVer) string {
	str := fmt.Sprintf("%u.%u.%u", ver.major, ver.minor, ver.patch)

	if len(ver.prerelease) > 0 {
		str += "-" + strings.Join(ver.prerelease, ".")
	}

	if len(ver.metadata) > 0 {
		str += "+" + strings.Join(ver.metadata, ".")
	}

	return str
}
