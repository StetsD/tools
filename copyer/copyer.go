package main

import (
	"flag"
	"fmt"
	"log"
)

var src, dest string

func init() {
	flag.StringVar(&src, "src", "", "path from")
	flag.StringVar(&dest, "dest", "./", "path to")
	flag.Parse()
}

func copy(src, dest string) (int64, error) {
	_, err := validator(src, dest)

	if err != nil {
		return 0, err
	}

	return 0, nil
}

func main() {
	if src == "" {
		log.Fatal(&Error{"\"src\" flag must be defined\n"})
	}

	_, err := copy(src, dest)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(src, dest)
}

// copyer -src /path/src -dest /path/dest -offset 1024 -limit 2048

// dest is dir
