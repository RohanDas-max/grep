package utils

import "fmt"

func PrintLine(data []string) {
	for i := range data {
		if data[i] == "" {
			i++
		} else {
			fmt.Println(data[i])
		}
	}
}
