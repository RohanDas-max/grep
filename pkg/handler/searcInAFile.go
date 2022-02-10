package handler

import (
	"github.com/rohandas-max/grep/pkg/utils"
)

func SearchInAFile(f, a string, c map[string]bool, cf map[string]int) ([]string, int, error) {
	data, err := utils.ReadFile(f)
	// fmt.Println(data)
	if err != nil {
		return []string{}, 0, err
	} else {
		s, _, count := utils.Search(data, a, "", false, c, cf)
		return s, count, nil
	}
}
