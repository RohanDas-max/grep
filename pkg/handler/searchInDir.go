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
			res = utils.Search(data, a, path, res)
			return nil
		}
	}); err != nil {
		return res, err
	}

	return res, nil
}
