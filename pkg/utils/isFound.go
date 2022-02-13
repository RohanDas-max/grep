package utils

import "strings"

func isFound(txt, arg string) (string, bool) {
	if strings.Contains(txt, arg) {
		return txt, true
	} else {
		return "", false
	}
}
