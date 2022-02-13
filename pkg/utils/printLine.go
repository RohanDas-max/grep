package utils

import (
	"fmt"
)

func PrintLine(data []string) {
	for _, line := range data {
		if line != "" {
			fmt.Printf("%v\n", line)
		}
	}
}
