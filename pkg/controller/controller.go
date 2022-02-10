package controller

import (
	"errors"
	"os"

	"github.com/rohandas-max/grep/pkg/handler"
)

func Controller(args []string, c map[string]bool, cf map[string]int) ([]string, int, error) {
	res := []string{}
	switch len(args) {
	case 1:
		return res, 0, handler.SearchInStdin(args[0], c, cf)
	case 2:
		fs, err := os.Stat(args[1])
		if err != nil {
			return res, 0, err
		}
		if fs.IsDir() {
			return handler.SearchInDir(args[1], args[0], c, cf)
		} else if !fs.IsDir() {
			return handler.SearchInAFile(string(args[1]), args[0], c, cf)
		}
	default:
		return res, 0, errors.New("oops! wrong command")
	}
	return res, 0, nil
}
