package reusablemethods

import "strings"

func IsMatched(text, arg string) string {
	if strings.Contains(text, arg) {
		return text
	} else {
		return ""
	}
}
