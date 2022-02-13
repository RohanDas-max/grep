package main

import (
	"flag"
	"learn/pkg/controller"
	"learn/pkg/utils"
)

func main() {

	i := flag.Bool("i", false, "-i option to search in case insensitive manner")
	c := flag.Bool("c", false, "-c option to count the result")
	flag.Parse()
	arg := flag.Arg(0)
	file := flag.Arg(1)
	option := make(map[string]bool)
	option["i"] = *i
	option["c"] = *c
	s := controller.Controller(file, arg, option)
	utils.PrintLine(s)

}
