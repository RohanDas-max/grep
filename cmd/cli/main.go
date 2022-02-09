package main

import (
	"flag"
	"log"

	"github.com/rohandas-max/grep/pkg/controller"
)

func main() {
	flag.Parse()
	args := flag.Args()
	err := controller.Controller(args)
	if err != nil {
		log.Fatal(err)
	}

}
