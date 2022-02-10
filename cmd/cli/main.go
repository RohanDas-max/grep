package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rohandas-max/grep/pkg/controller"
	"github.com/rohandas-max/grep/pkg/utils"
)

func main() {
	i := flag.Bool("i", false, "option to search in case-insensitive manner")
	c := flag.Bool("c", false, "option to count")
	C := flag.Int("C", 0, "")
	A := flag.Int("A", 0, "")
	B := flag.Int("B", 0, "")
	flag.Parse()
	args := flag.Args()
	var ControlFlag = make(map[string]bool)
	ControlFlag["i"] = *i
	ControlFlag["c"] = *c
	var ContextFlag = make(map[string]int)
	ContextFlag["A"] = *A
	ContextFlag["B"] = *B
	ContextFlag["C"] = *C

	res, count, err := controller.Controller(args, ControlFlag, ContextFlag)
	if err != nil {
		log.Fatal(err)
	} else if count > 0 {
		fmt.Printf("%d\n", count)
	} else {
		utils.PrintLine(res)
	}
}
