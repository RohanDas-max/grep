package handler

import (
	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInAFile(f, a string) ([]string, error) {
	data, err := utils.ReadFile(f)
	if err != nil {
		return []string{""}, err
	}
	var res []string
	for _, dt := range data {
		res = append(res, utils.IsMatched(dt, a))
	}

	return res, nil
}
