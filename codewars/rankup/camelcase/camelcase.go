package camelcase

import (
	"unicode"
)

func CamelCase(s string) string {
	if len(s) == 0 {
		return s
	}

	nextCapitalized := true

	var res string
	for _, r := range s {
		if r == ' ' {
			nextCapitalized = true
			continue
		}

		if nextCapitalized {
			res += string(unicode.ToUpper(r))
			nextCapitalized = false
		} else {
			res += string(r)
		}
	}

	return res
}
