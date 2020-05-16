package copyer

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func concatName(filePath string) (result string) {
	fileBase := filepath.Base(filePath)
	ext := filepath.Ext(fileBase)
	fileBase = strings.TrimSuffix(fileBase, ext)

	alreadyInc, err := regexp.Match(`(_\d+)$`, []byte(fileBase))

	if err != nil {
		panic(err)
	}

	if !alreadyInc {
		result = fileBase + "_01" + ext
	} else {
		var formattedInc string
		reg := regexp.MustCompile(`(_\d+)$`)
		inc := fmt.Sprintf("%s", reg.Find([]byte(fileBase)))
		fileBase = strings.TrimSuffix(fileBase, inc)
		inc = strings.TrimPrefix(inc, "_")
		incInt, err := strconv.Atoi(inc)

		if err != nil {
			panic(err)
		}

		if incInt <= 8 {
			formattedInc = "0" + strconv.Itoa(incInt+1)
		} else {
			formattedInc = strconv.Itoa(incInt + 1)
		}

		result = fileBase + "_" + formattedInc + ext
	}

	return result
}

func setIncrementName(dest, filePath string) (string, error) {
	if _, err := os.Stat(path.Join(dest, filePath)); err != nil {
		if os.IsNotExist(err) {
			return path.Join(dest, filePath), nil
		} else {
			return "", err
		}
	}

	return setIncrementName(dest, concatName(filePath))
}

func NameChecker(dest, filePath string) (string, error) {
	checkedPath, err := setIncrementName(dest, filepath.Base(filePath))

	if err != nil {
		return "", err
	}

	return checkedPath, nil
}
