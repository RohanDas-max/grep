package handler

import (
	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInAFile(f, a string) ([]string, error) {
	var res []string
	data, err := utils.ReadFile(f)
	if err != nil {
		return []string{}, err
	} else {
		return utils.Search(data, a, "", res), nil
	}
}
