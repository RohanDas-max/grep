package controller

import (
	"flag"
	"fmt"

	"github.com/rohandas-max/grep/pkg/handler"
)

func Controller() error {
	flag.Parse()
	arg := flag.Arg(0)
	fn := flag.Arg(1)
	data, err := handler.SearchInAFile(fn, arg)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", data)
	return nil
}
