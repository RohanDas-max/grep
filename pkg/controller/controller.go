package controller

import (
	"errors"
	"fmt"
	"os"

	"github.com/rohandas-max/grep/pkg/handler"
	"github.com/rohandas-max/grep/pkg/utils"
)

func Controller(args []string) error {
	switch {
	case len(args) <= 1:
		if err := handler.SearchInStdin(string(args[0])); err != nil {
			return err
		}
	case len(args) <= 2:
		fs, err := os.Stat(args[1])
		if err != nil {
			return err
		}
		if fs.IsDir() {
			data, err := handler.SearchInDir(args[1])
			if err != nil {
				return errors.New("invalid folder name")
			}
			utils.PrintLine(data, 0)
		} else if !fs.IsDir() {
			data, err := handler.SearchInAFile(string(args[1]), string(args[0]))
			if err != nil {
				return errors.New("invalid file name")

			}
			utils.PrintLine(data, 1)
		}
	default:
		fmt.Println("oops! wrong command")
	}
	return nil
}
