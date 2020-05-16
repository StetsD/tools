package copyer

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clearCmd func()

func init() {
	clearMap := make(map[string]func())
	clearMap["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	clearMap["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	val, ok := clearMap[runtime.GOOS]

	if ok {
		clearCmd = val
	} else {
		log.Fatal("spinner can not correct inits, platform not defined")
	}
}

func spinner(spinPosition chan<- rune, duration time.Duration, done <-chan bool) {
	for {
		if _, ok := <-done; ok {
			break
		}
		for _, v := range "-\\|/" {
			spinPosition <- v
			time.Sleep(duration)
		}
	}
}

func StatSpinner(writtenBytes <-chan int, totalBytes float64, done <-chan bool) {
	spinPosition := make(chan rune)
	restBytes := totalBytes

	go spinner(spinPosition, time.Microsecond*200, done)

	for {
		select {

		case spinnerPos := <-spinPosition:
			clearCmd()
			fmt.Printf("\r%c", spinnerPos)

		case bytes := <-writtenBytes:
			{
				clearCmd()
				restBytes -= float64(bytes)
				progress := math.Ceil(100 - ((restBytes / totalBytes) * 100))
				fmt.Printf("\tTotal bytes written %d\n\tProgress %v%%", bytes, progress)
			}

		case _ = <-done:
			break

		}
	}

}
