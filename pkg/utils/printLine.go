package utils

import "fmt"

func PrintLine(s []string, signal int) {
	// sp := "\n"
	for i := range s {
		if s[i] == "" {
			i++
		} else if signal > 0 {
			fmt.Printf("%v\n", s[i])
		} else {
			fmt.Printf("%v", s[i])

		}
	}
}
