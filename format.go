package semver

import (
	"fmt"
	"strings"
)

// String converts the semver argument into a string of the form
// <major>.<minor>.<patch>[-<pre.release>][+<meta.data>]
func String(ver SemVer) string {
	str := fmt.Sprintf("%u.%u.%u", ver.Major, ver.Minor, ver.Patch)

	if len(ver.Prerelease) > 0 {
		str += "-" + strings.Join(ver.Prerelease, ".")
	}

	if len(ver.Metadata) > 0 {
		str += "+" + strings.Join(ver.Metadata, ".")
	}

	return str
}
