package handler

import (
	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInAFile(f, a string) ([]string, error) {
	data, err := utils.ReadFile(f)
	// fmt.Println(data)
	if err != nil {
		return []string{}, err
	} else {
		s, _ := utils.Search(data, a, "", false)
		return s, nil
	}
}
