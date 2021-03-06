package copyer

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

var wd string

func init() {
	_, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	wd, _ = os.Getwd()
}

func Copy(src, dest string, bytesLimit int) (int64, error) {
	var totalSizeOfBytes int64
	writtenBytes := make(chan int)
	doneChan := make(chan bool)
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
		if bytesLimit != 0 {
			totalSizeOfBytes += int64(bytesLimit)
		} else {
			totalSizeOfBytes += totalSizeOfBytes + srcPath.size
		}
	}

	go StatSpinner(writtenBytes, float64(totalSizeOfBytes), doneChan)

	for _, srcPath := range srcCollection {
		fileSrc, _ := os.Open(srcPath.path)
		buf := make([]byte, srcPath.size)
		offset := 0
		fileDestPath, err := NameChecker(dest, srcPath.path)

		if err != nil {
			err := fmt.Errorf("%s\n", err)
			fmt.Println(err.Error())
			continue
		}

		fileDest, err := os.Create(fileDestPath)

		if err != nil {
			err := fmt.Errorf("%s\n", err)
			fmt.Println(err.Error())
			continue
		}

	toBuf:
		for {
			if bytesLimit != 0 && bytesLimit <= offset {
				break toBuf
			}

			read, err := fileSrc.Read(buf[offset:])

			if err == io.EOF {
				break toBuf
			}

			if read == 0 {
				break toBuf
			}

			if err != nil {
				log.Fatal(err)
			}

			if bytesLimit != 0 && bytesLimit < read {
				read = bytesLimit
			}

			if _, err := fileDest.Write(buf[:read]); err != nil {
				log.Fatal(err)
			}

			writtenBytes <- read

			offset += read
		}

		err = fileSrc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	doneChan <- true
	close(writtenBytes)
	return 0, nil
}
