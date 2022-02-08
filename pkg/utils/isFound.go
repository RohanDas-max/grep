package utils

import (
	"strings"
)

func IsFound(txt, a string) string {
	if strings.Contains(txt, a) {
		return txt
	}
	return ""
}
