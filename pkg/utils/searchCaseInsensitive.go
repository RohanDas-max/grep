package utils

import "fmt"

func SearchCaseInsensitive(data []string, a, path string, stdin bool) ([]string, string) {
	res := []string{}
	var s string
	if !stdin {
		for _, line := range data {
			if v, found := isFoundCaseIns(line, a); found {
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
		if v, found := isFoundCaseIns(data[0], a); found {
			return res, v
		}
	}

	return res, s
}
