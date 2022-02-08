package handler

import (
	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInAFile(f, a string) ([]string, error) {

	datas, err := utils.ReadFile(f)
	if err != nil {
		return []string{""}, err
	}
	var res []string
	for _, data := range datas {
		res = append(res, utils.IsMatched(data, a))
	}

	return res, nil
}
