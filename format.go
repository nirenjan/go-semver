package semver

import (
	"fmt"
	"strings"
)

// String converts the semver argument into a string of the form
// <major>.<minor>.<patch>[-<pre.release>][+<meta.data>]
func (ver SemVer) String() string {
	str := fmt.Sprintf("%u.%u.%u", ver.Major, ver.Minor, ver.Patch)

	if len(ver.Prerelease) > 0 {
		str += "-" + strings.Join(ver.Prerelease, ".")
	}

	if len(ver.Metadata) > 0 {
		str += "+" + strings.Join(ver.Metadata, ".")
	}

	return str
}

// Tag converts the semver argument into a string of the form
// v<major>.<minor>.<patch>[-<pre.release>][+<meta.data>]
// This is suitable for inserting as a tag into your version control system
func (ver SemVer) Tag() string {
	return "v" + ver.String()
}
