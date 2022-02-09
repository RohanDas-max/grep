package main

import (
	"flag"
	"log"

	"github.com/rohandas-max/grep/pkg/controller"
)

func main() {
	i := flag.Bool("i", false, "option to search in case-insensitive manner")
	flag.Parse()
	args := flag.Args()
	err := controller.Controller(args, *i)
	if err != nil {
		log.Fatal(err)
	}

}
