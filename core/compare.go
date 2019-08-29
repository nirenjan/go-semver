package core

import "strconv"

// Compare compares s1 against s2, and returns the following values
// negative if s1 is less than s2
// 0 if s1 and s2 are identical
// positive if s1 is greater than s2
func (s1 SemVer) Compare(s2 SemVer) int {
	if s1.major != s2.major {
		return spaceShip((int)(s1.major - s2.major))
	}

	if s1.minor != s2.minor {
		return spaceShip((int)(s1.minor - s2.minor))
	}

	if s1.patch != s2.patch {
		return spaceShip((int)(s1.patch - s2.patch))
	}

	return comparePrereleases(s1.prerelease, s2.prerelease)
}

// comparePrereleases checks the prereleases according to the following rules:
// 1. An empty prerelease ranks higher than a non-empty prerelease
// 2. Prerelease identifiers are compared by individual dot-separated identifiers.
// 3. A numeric identifier ranks lower than a non-numeric identifier
// 4. A prerelease string that has more identifiers ranks higher than one with
//    fewer identifiers
func comparePrereleases(s1, s2 []string) int {
	if len(s1) == 0 || len(s2) == 0 {
		return spaceShip(len(s2) - len(s1))
	}

	// Iterate over the shorter of the two arrays
	var count int
	if len(s1) < len(s2) {
		count = len(s1)
	} else {
		count = len(s2)
	}

	for i := 0; i < count; i++ {
		if ret := compareIdentifiers(s1[i], s2[i]); ret != 0 {
			return ret
		}
	}

	// At this point, we've gone over the common length of the arrays
	// and they are identical so far. So the longer array ranks higher
	return spaceShip(len(s1) - len(s2))
}

// compareIdentifiers checks the identifiers and ranks them accordingly.
func compareIdentifiers(i1, i2 string) int {
	// Try to convert the strings to the integer values
	n1, err1 := strconv.Atoi(i1)
	n2, err2 := strconv.Atoi(i2)

	if err1 == nil {
		// i1 is numeric
		if err2 == nil {
			// Both could be converted to int
			return spaceShip(n1 - n2)
		} else {
			// i1 is numeric, but i2 is not.
			// Return -1 since numeric compares lower than non-numeric
			return -1
		}
	} else {
		// i1 is non-numeric
		if err2 == nil {
			// i2 is numeric, but i1 is not.
			// Return +1 since non-numeric compares higher than numeric
			return +1
		}
	}

	// At this point, both are non-numeric. We need to determine the
	// return value based on string comparision
	if i1 == i2 {
		return 0
	}

	if i1 < i2 {
		return -1
	}

	return +1
}

// spaceShip converts negative numbers to -1, and positive numbers to +1
// This is to make the three-way comparison consistent
func spaceShip(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return +1
	}

	return 0
}
