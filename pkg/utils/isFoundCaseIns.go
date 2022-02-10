package utils

import (
	"strings"
)

func isFoundCaseIns(txt, a string) (string, bool) {
	if strings.Contains(strings.ToLower(txt), strings.ToLower(a)) {
		return txt, true
	} else {
		return "", false
	}
}
