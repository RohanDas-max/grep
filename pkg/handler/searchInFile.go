package handler

import "learn/pkg/utils"

func SearchInFile(name, arg string) []string {
	d := utils.Readfile(name)
	res := utils.Search(d, arg)
	return res
}

func SearchInFileCaseIns(name, arg string) []string {
	d := utils.Readfile(name)
	res := utils.SearchCaseInsensitive(d, arg)
	return res
}
