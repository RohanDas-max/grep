package handler

import (
	"io/fs"
	"path/filepath"

	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInDir(f, a string, c map[string]bool, cf map[string]int) ([]string, int, error) {
	var res = []string{}
	var num int
	if err := filepath.Walk(f, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		} else {
			var s []string
			var count int
			data, _ := utils.ReadFile(path)
			s, _, count = utils.Search(data, a, path, false, c, cf)
			num += count
			res = append(res, s...)
			return nil
		}
	}); err != nil {
		return []string{}, 0, err
	}

	return res, num, nil
}
