package controller

import (
	"fmt"

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
		data, err := handler.SearchInAFile(string(args[1]), string(args[0]))
		if err != nil {
			return err
		}
		utils.PrintLine(data)
	default:
		fmt.Println("oops! wrong input")
	}
	return nil
}
