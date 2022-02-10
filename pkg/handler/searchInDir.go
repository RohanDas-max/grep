package handler

import (
	"io/fs"
	"path/filepath"

	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInDir(f, a string) ([]string, error) {
	var res []string

	if err := filepath.Walk(f, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		} else {
			data, _ := utils.ReadFile(path)
			// fmt.Println(path, data)
			s, _ := utils.Search(data, a, path, false)
			// fmt.Println(s)
			res = append(res, s...)

			return nil
		}
	}); err != nil {
		return res, err
	}

	return res, nil
}
