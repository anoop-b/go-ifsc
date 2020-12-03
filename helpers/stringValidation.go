package helpers

import "strings"

func ValidIfsc(ifsc string) (string, bool) {
	if len(ifsc) != 11 {
		return "", false
	}
	return strings.ToUpper(ifsc), true
}
