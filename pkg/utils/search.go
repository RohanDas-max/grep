package utils

import "fmt"

func Search(data []string, a, path string, res []string) []string {
	for _, line := range data {
		if v, found := IsFound(line, a); found {
			if path != "" {
				s := fmt.Sprintf("%v:%v", path, v)
				res = append(res, s)
			} else {
				res = append(res, v)
			}
		}
	}
	return res
}
