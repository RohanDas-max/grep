package utils

func SearchCaseInsensitive(data []string, arg string) []string {
	var result []string
	for _, line := range data {
		if v, found := isFoundCaseIns(line, arg); found {
			result = append(result, v)
		}
	}
	return result
}
