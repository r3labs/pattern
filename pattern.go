package pattern

import "strings"

// Match : will match a string pattern using wildcards
func Match(value string, params ...string) bool {
	for _, p := range params {
		if patternMatches(value, p) || p == value {
			return true
		}
	}

	return false
}

func patternMatches(value, match string) bool {
	if !strings.Contains(match, "*") {
		return false
	}

	sparts := strings.Split(value, ".")
	mparts := strings.Split(match, ".")

	if len(sparts) != len(mparts) {
		return false
	}

	for i, m := range mparts {
		if m == "*" {
			continue
		}

		if sparts[i] != m {
			return false
		}
	}

	return true
}
