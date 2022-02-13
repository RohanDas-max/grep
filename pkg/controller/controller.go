package controller

import (
	"io/fs"
	"learn/pkg/handler"
	"os"
	"strconv"
)

func Controller(filename, arg string, option map[string]bool) []string {
	fs, _ := os.Stat(filename)
	if option["i"] && option["c"] {
		s := caseInsensitive(filename, arg, fs, option["c"])
		if len(s) != 0 {
			return []string{strconv.Itoa(len(s))}
		}
	} else if option["i"] {
		return caseInsensitive(filename, arg, fs, option["c"])
	} else if option["c"] {
		s := caseSensitive(filename, arg, fs, option["c"])
		if len(s) != 0 {
			return []string{strconv.Itoa(len(s))}
		}
	} else {
		return caseSensitive(filename, arg, fs, option["c"])
	}
	return []string{}
}

func caseSensitive(filename, arg string, fs fs.FileInfo, c bool) []string {
	if filename == "" {
		handler.SearcInStdinCaseIns(arg, c)
	} else if fs.IsDir() {
		return handler.SearcInDir(filename, arg)
	} else if !fs.IsDir() {
		return handler.SearchInFile(filename, arg)
	}
	return []string{}
}

func caseInsensitive(filename, arg string, fs fs.FileInfo, c bool) []string {
	if filename == "" {
		handler.SearcInStdinCaseIns(arg, c)
	} else if fs.IsDir() {
		return handler.SearchInDirCaseIns(filename, arg)
	} else if !fs.IsDir() {
		return handler.SearchInFileCaseIns(filename, arg)
	}
	return []string{}
}
