package utils

import (
	"os"
	"strings"
)

func ReadFile(f string) ([]string, error) {
	dataByte, err := os.ReadFile(f)
	if err != nil {
		return []string{}, err
	} else {
		line := strings.Split(string(dataByte), "\n")
		return line, nil
	}
}
