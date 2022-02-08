package main

import (
	"flag"
	"fmt"

	"github.com/rohandas-max/grep/pkg/controller"
)

func main() {
	flag.Parse()
	args := flag.Args()
	err := controller.Controller(args)
	if err != nil {
		panic(err)
	}
	func() {
		active := recover()
		fmt.Printf("%v", active)
	}()
}
