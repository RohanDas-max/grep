package controller

import (
	"errors"
	"os"
	"strings"

	"github.com/rohandas-max/grep/pkg/handler"
	"github.com/rohandas-max/grep/pkg/utils"
)

func Controller(args []string, i bool) error {
	a := string(args[0])
	if i {
		a = strings.ToLower(args[0])
	}
	switch {
	case len(args) < 1:
		return errors.New("please add some option or command")
	case len(args) <= 1:
		if err := handler.SearchInStdin(a); err != nil {
			return err
		}
	case len(args) <= 2:
		fs, err := os.Stat(args[1])
		if err != nil {
			return err
		}
		if fs.IsDir() {
			data, err := handler.SearchInDir(args[1], a)
			if err != nil {
				return errors.New("invalid folder name")
			}
			utils.PrintLine(data)
		} else if !fs.IsDir() {
			data, err := handler.SearchInAFile(string(args[1]), a)
			if err != nil {
				return errors.New("invalid file name")

			}
			utils.PrintLine(data)
		}
	default:
		return errors.New("oops! wrong command")
	}
	return nil
}
