package handler

import (
	"bufio"
	"fmt"
	"io"
	"learn/pkg/utils"
	"log"
	"os"
)

func SearcInStdin(arg string, c bool) {
	reader := bufio.NewReader(os.Stdin)
	var num int
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			if c {
				fmt.Println(num)
			}
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			var data = []string{s}
			res := utils.Search(data, arg)
			if res != nil && !c {
				if !c {
					fmt.Printf("%v", res[0])
				}
				num++
			}
		}
	}

}

func SearcInStdinCaseIns(arg string, c bool) {

	reader := bufio.NewReader(os.Stdin)
	var num int
	for {
		s, err := reader.ReadString('\n')
		if err == io.EOF {
			if c {
				fmt.Println(num)
			}
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			var data = []string{s}
			res := utils.SearchCaseInsensitive(data, arg)
			if res != nil {
				if !c {
					fmt.Printf("%v", res[0])
				}
				num++
			}
		}
	}

}
