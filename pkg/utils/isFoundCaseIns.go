package utils

import "strings"

func isFoundCaseIns(txt, arg string) (string, bool) {
	if strings.Contains(strings.ToLower(txt), strings.ToLower(arg)) {
		return txt, true
	} else {
		return "", false
	}
}
