package utils

func Search(data []string, a, path string, stdin bool, c map[string]bool, cf map[string]int) ([]string, string, int) {
	cInsesitive := c["i"]
	cCount := c["c"]

	// cFA := cf["A"]
	// cFB := cf["B"]
	// cFC := cf["C"]

	if cInsesitive && cCount {
		reS, s := SearchCaseInsensitive(data, a, path, stdin)
		return []string{}, "", count(reS, s)
	} else if cInsesitive {
		reS, s := SearchCaseInsensitive(data, a, path, stdin)
		return reS, s, 0
	} else if cCount {
		reS, s := SearchCaseSensitive(data, a, path, stdin)
		return []string{}, "", count(reS, s)
	} else {
		reS, s := SearchCaseSensitive(data, a, path, stdin)
		return reS, s, 0
	}
}
