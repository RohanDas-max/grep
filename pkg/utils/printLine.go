package utils

import (
	"errors"
	"fmt"
)

func PrintLine(s []string) error {
	if len(s) < 1 {
		return errors.New("slice is empty nothing to print")
	} else {
		for i := range s {
			if s[i] == "" {
				i++
			} else {
				fmt.Printf("%v\n", s[i])
			}
		}
		return nil
	}
}
