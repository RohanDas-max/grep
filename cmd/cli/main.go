package main

import (
	"flag"
	"log"

	"github.com/rohandas-max/grep/pkg/controller"
	"github.com/rohandas-max/grep/pkg/utils"
)

func main() {
	i := flag.Bool("i", false, "option to search in case-insensitive manner")
	flag.Parse()
	args := flag.Args()
	res, err := controller.Controller(args, *i)
	if err != nil {
		log.Fatal(err)
	}
	utils.PrintLine(res)

}
