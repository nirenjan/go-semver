// Test routines for package semver

package semver

import "testing"

// TestString checks the output of semver.String
func TestString(t *testing.T) {
	var tests = []struct {
		major      uint
		minor      uint
		patch      uint
		prerelease string
		metadata   string
		expected   string
	}{
		{0, 1, 0, "", "", "0.1.0"},
		{1, 0, 0, "", "", "1.0.0"},
		{1, 2, 3, "", "", "1.2.3"},
		{0, 12, 0, "", "", "0.12.0"},
		{0, 1, 20, "", "", "0.1.20"},
		{0, 2, 3, "alpha.1", "", "0.2.3-alpha.1"},
		{0, 2, 3, "beta", "", "0.2.3-beta"},
		{0, 2, 3, "", "gabcdef0", "0.2.3+gabcdef0"},
		{0, 2, 3, "", "gabcdef0.2015-04-20T1001-0700", "0.2.3+gabcdef0.2015-04-20T1001-0700"},
		{0, 2, 3, "beta", "20150420", "0.2.3-beta+20150420"},
		{0, 2, 3, "rc.1", "2015.04.20", "0.2.3-rc.1+2015.04.20"},
	}

	for _, test := range tests {
		var ver SemVer

		ver.Major = test.major
		ver.Minor = test.minor
		ver.Patch = test.patch
		ver.Prerelease(test.prerelease)
		ver.Metadata(test.metadata)

		got := ver.String()

		if got != test.expected {
			t.Errorf("(%v.%v.%v pre %q meta %q) expected %q, got %q",
				test.major, test.minor, test.patch,
				test.prerelease, test.metadata, test.expected, got)
		}
	}
}

// TestTag checks the output of semver.Tag
func TestTag(t *testing.T) {
	var tests = []struct {
		major      uint
		minor      uint
		patch      uint
		prerelease string
		metadata   string
		expected   string
	}{
		{0, 1, 0, "", "", "v0.1.0"},
		{0, 2, 3, "beta", "", "v0.2.3-beta"},
		{0, 2, 3, "", "gabcdef0", "v0.2.3+gabcdef0"},
		{0, 2, 3, "rc.1", "2015.04.20", "v0.2.3-rc.1+2015.04.20"},
	}

	for _, test := range tests {
		var ver SemVer

		ver.Major = test.major
		ver.Minor = test.minor
		ver.Patch = test.patch
		ver.Prerelease(test.prerelease)
		ver.Metadata(test.metadata)

		got := ver.Tag()

		if got != test.expected {
			t.Errorf("(%v.%v.%v pre %q meta %q) expected %q, got %q",
				test.major, test.minor, test.patch,
				test.prerelease, test.metadata, test.expected, got)
		}
	}
}
