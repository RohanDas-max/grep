package utils

import (
	"os"
	"strings"
)

func Readfile(f string) []string {
	s, _ := os.ReadFile(f)
	return strings.Split(string(s), "\n")
}
