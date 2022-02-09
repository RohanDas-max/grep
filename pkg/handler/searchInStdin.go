package handler

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInStdin(arg string) error {
	if arg != "" {
		reader := bufio.NewReader(os.Stdin)
		for {
			text, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				return err
			} else {
				res := utils.IsFound(text, arg)
				fmt.Printf("%s", res)
			}
		}
	} else {
		return errors.New("empty argument")
	}
	return nil
}