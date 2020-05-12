package main

import (
	"io/ioutil"
	"os"
	"path"
)

func collect(src string, collection *[]string) (err error) {
	fi, err := os.Stat(src)

	if err != nil {
		return err
	}

	switch mode := fi.Mode(); {

	case mode.IsDir():
		files, err := ioutil.ReadDir(src)

		if err != nil {
			return err
		}

		for _, file := range files {
			err := collect(path.Join(src, file.Name()), collection)

			if err != nil {
				return err
			}
		}

	case mode.IsRegular():
		*collection = append(*collection, src)
	default:
		panic(&Error{"src path must be include only regular files and directories"})

	}

	return nil
}

func Collector(src string) (result []string, err error) {
	err = collect(src, &result)

	if err != nil {
		return result, err
	}

	return result, nil
}
