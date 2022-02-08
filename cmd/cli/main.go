package main

import (
	"flag"

	"github.com/rohandas-max/grep/pkg/controller"
)

func main() {
	flag.Parse()
	args := flag.Args()
	controller.Controller(args)
}
