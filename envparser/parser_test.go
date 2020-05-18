package envparser

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"
	"testing"
)

var envMap = map[string]string{
	"COUNTY":   "Afghanistan",
	"COMPANY":  "Taliban",
	"PASSWORD": "HollyDiver",
}

func TestAppliedEnv(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filePath := path.Join(wd, ".env")
	testEnvFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := testEnvFile.Close()
		if err != nil {
			log.Fatal(err)
		}
		err = os.Remove(filePath)
		if err != nil {
			log.Fatal(err)
		}
	}()

	writer := bufio.NewWriter(testEnvFile)
	for key, val := range envMap {
		_, err := writer.WriteString(key + "=" + val + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}

	oldEnv := os.Environ()

	parsedEnv, err := Parse(filePath)
	if err != nil {
		log.Fatal(err)
	}
	Apply(parsedEnv)

	newEnv := os.Environ()

	if len(newEnv) != len(oldEnv)+len(envMap) {
		t.Error("Some variables did not apply")
	}

outer:
	for key, val := range envMap {
		for i, v := range newEnv {
			splited := strings.Split(v, "=")

			if splited[0] == key && splited[1] == val {
				continue outer
			}

			if len(newEnv)-1 == i {
				if splited[0] != key && splited[1] != val {
					t.Error("env \"" + key + "\" did not apply")
				}
			}
		}
	}
}
