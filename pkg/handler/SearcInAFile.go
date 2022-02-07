package handler

import (
	"os"

	reusablemethods "github.com/rohandas-max/grep/pkg/reusableMethods"
)

func SearchInAFile(filename, arg string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return reusablemethods.IsMatched(string(data), arg), nil
}
