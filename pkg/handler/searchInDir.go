package handler

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInDir(f string) ([]string, error) {
	var res []string

	if err := filepath.Walk(f, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		} else {
			data, _ := utils.ReadFile(path)
			for _, data := range data {
				if data != "" {
					s := fmt.Sprintln(path, ":", utils.IsFound(data, "hello"))
					res = append(res, s)
				}
			}
			return nil
		}
	}); err != nil {
		return []string{""}, err
	}
	return res, nil
}
