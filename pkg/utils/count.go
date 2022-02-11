package utils

func count(reS []string, s string) int {
	var count int
	if s == "" {
		return len(reS)
	} else {
		count += 1
	}
	return count
}
