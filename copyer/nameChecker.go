package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func concatName(path string) string {
	fileBase := filepath.Base(path)
	ext := filepath.Ext(fileBase)
	fileBase = strings.TrimSuffix(fileBase, ext)

	alreadyInc, err := regexp.Match(`(_\d+)$`, []byte(fileBase))

	if err != nil {
		panic(err)
	}

	if !alreadyInc {
		return fileBase + "_01" + ext
	} else {
		var formattedInc string
		reg := regexp.MustCompile(`(_\d+)$`)
		inc := fmt.Sprintf("%s", reg.Find([]byte(fileBase)))
		inc = strings.TrimPrefix(inc, "_")
		incInt, err := strconv.Atoi(inc)

		if err != nil {
			panic(err)
		}

		if incInt <= 9 {
			formattedInc = "0" + strconv.Itoa(incInt+1)
		} else {
			formattedInc = strconv.Itoa(incInt + 1)
		}

		return fileBase + "_" + formattedInc + ext
	}
}

func setIncrementName(path string) (string, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return path, nil
		} else {
			return "", err
		}
	}

	return setIncrementName(concatName(path))
}

func NameChecker(path string) (string, error) {
	checkedPath, err := setIncrementName(path)

	if err != nil {
		return "", err
	}

	return checkedPath, nil
}
