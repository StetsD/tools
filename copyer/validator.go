package copyer

import (
	"log"
	"os"
)

func validator(src, dest string) (int64, error) {
	pathMap := map[string]string{
		"src":  src,
		"dest": dest,
	}

	for k, v := range pathMap {
		fi, err := os.Stat(v)

		if err != nil {
			return 0, err
		}

		if k == "dest" && !fi.IsDir() {
			log.Fatal(&Error{"dest must be a directory"})
		}

	}

	return 0, nil
}
