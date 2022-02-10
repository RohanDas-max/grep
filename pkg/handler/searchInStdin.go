package handler

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInStdin(arg string, c map[string]bool, cf map[string]int) error {

	if arg != "" {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			} else {
				var data = []string{text}
				_, res, count := (utils.Search(data, arg, "", true, c, cf))
				if count > 0 {
					fmt.Printf("%d", count)
				} else {
					fmt.Printf("%s", res)
				}
			}
		}
	} else {
		return errors.New("empty argument")
	}
	return nil
}
