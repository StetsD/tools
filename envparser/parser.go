package envparser

import (
	"bufio"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

var wd string

func init() {
	wd, _ = os.Getwd()
}

func Parse(filepath string) (map[string]string, error) {
	var envMap = make(map[string]string)

	if !path.IsAbs(filepath) {
		filepath = path.Join(wd, filepath)
	}

	file, err := os.Open(filepath)
	if err != nil {
		return envMap, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		isValidEnv, err := regexp.Match(`([A-Z]|_)+=\w+`, []byte(scanner.Text()))
		if err != nil {
			panic(err)
		}

		if isValidEnv {
			splitted := strings.Split(scanner.Text(), "=")
			envMap[splitted[0]] = splitted[1]
		}
	}

	return envMap, nil
}

func Apply(envs map[string]string) {
	for key, val := range envs {
		err := os.Setenv(key, val)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
