package handler

import (
	"fmt"
	"io/fs"
	"learn/pkg/utils"
	"log"
	"path/filepath"
)

func SearcInDir(name, arg string) []string {
	var result = []string{}
	if err := filepath.Walk(name, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		data := utils.Readfile(path)

		var res []string
		if v := utils.Search(data, arg); v != nil {
			res = append(res, fmt.Sprint(path, ":", v))
		}
		result = append(result, res...)

		return nil
	}); err != nil {
		log.Fatal("in search in dir", err)
	}

	return result
}

func SearchInDirCaseIns(name, arg string) []string {
	var result = []string{}
	if err := filepath.Walk(name, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		data := utils.Readfile(path)

		var res []string
		if v := utils.SearchCaseInsensitive(data, arg); v != nil {
			res = append(res, fmt.Sprint(path, ":", v))
		}
		result = append(result, res...)

		return nil
	}); err != nil {
		log.Fatal("in search in dir", err)
	}
	return result
}
