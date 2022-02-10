package utils

import "fmt"

func Search(data []string, a, path string, stdin bool) ([]string, string) {
	res := []string{}
	var s string
	if !stdin {
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
		return res, s
	} else if stdin {
		if v, found := IsFound(data[0], a); found {
			return res, v
		}
	}
	return res, s
}
