package controller

import (
	"flag"
	"fmt"

	"github.com/rohandas-max/grep/pkg/handler"
)

func Controller() error {
	flag.Parse()
	a := flag.Arg(0)
	f := flag.Arg(1)

	data, err := handler.SearchInAFile(f, a)
	if err != nil {
		return err
	}
	for i := range data {
		if data[i] == "" {
			i++
		} else {
			fmt.Println(data[i])
		}
	}
	return nil
}
