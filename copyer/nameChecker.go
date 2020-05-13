package main

import (
	"os"
)

func setIncrementName(path string) (string, error) {
	if _, err := os.Stat(path); err != nil {
		if err != nil {
			if os.IsNotExist(err) {
				return path, nil
			}
		}
	}

	return setIncrementName(path + "")
}

func NameChecker(path string) (string, error) {
	checkedPath, err := setIncrementName(path)

	if err != nil {
		return "", err
	}

	return checkedPath, nil
}
