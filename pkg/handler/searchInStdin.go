package handler

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInStdin(arg string) error {
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
	return nil
}
