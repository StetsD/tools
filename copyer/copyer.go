package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
)

var src, dest string
var bytesWindow uint64
var wd string

const defaultBytesWindow = 1024

func init() {
	_, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	wd, _ = os.Getwd()

	flag.StringVar(&src, "src", "", "path from")
	flag.StringVar(&dest, "dest", wd, "path to")
	flag.Uint64Var(&bytesWindow, "b", defaultBytesWindow, "window of bytes")
	flag.Parse()
}

func Copy(src, dest string) (int64, error) {

	_, err := validator(src, dest)
	if err != nil {
		return 0, err
	}

	srcCollection, err := Collector(src)
	if err != nil {
		return 0, err
	}

	if !path.IsAbs(dest) {
		dest = path.Join(wd, dest)
	}

	for _, srcPath := range srcCollection {
		fileSrc, _ := os.Open(srcPath)
		fileDest, err := NameChecker(dest, srcPath)

		if err != nil {
			fmt.Errorf("%s\n", err)
			continue
		}

	}

	return 0, nil
}

func main() {
	if src == "" {
		log.Fatal(&Error{"\"src\" flag must be defined\n"})
	}

	_, err := Copy(src, dest)

	if err != nil {
		log.Fatal(err)
	}
}
