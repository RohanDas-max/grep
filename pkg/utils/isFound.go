package utils

import (
	"strings"
)

func IsMatched(txt, a string) string {
	if strings.Contains(txt, a) {
		return txt
	}
	return ""
}
