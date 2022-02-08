package utils

import "fmt"

func PrintLine(s []string) {
	for i := range s {
		if s[i] == "" {
			i++
		} else {
			fmt.Println(s[i])
		}
	}
}
