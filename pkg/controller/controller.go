package controller

import (
	"errors"
	"os"
	"strings"

	"github.com/rohandas-max/grep/pkg/handler"
)

func Controller(args []string, iFlag bool) ([]string, error) {
	var res []string
	a := string(args[0])
	if iFlag {
		a = strings.ToLower(args[0])
	}

	switch len(args) {
	case 1:
		return res, handler.SearchInStdin(a)
	case 2:
		fs, err := os.Stat(args[1])
		if err != nil {
			return res, err
		}
		if fs.IsDir() {
			return handler.SearchInDir(args[1], a)
		} else if !fs.IsDir() {
			return handler.SearchInAFile(string(args[1]), a)
		}
	default:
		return res, errors.New("oops! wrong command")
	}
	return res, nil
}
