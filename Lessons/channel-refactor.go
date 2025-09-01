package main

import (
	"fmt"
	"log"
	"io"
	"os"
	"strings"
)

// the goal of this is to turn our state machine into a goroutine with a channel to carry out each line


func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string) // the channel that will carry lines out

	go func() { // goroutine
		defer close(out)
		defer f.Close()

		curLine := ""
		data := make([]byte, 8)

		for {
			n, err := f.Read(data)
			if err != nil {
				break
			}

			parts := strings.Split(string(data[:n]), "\n")		

			for i := 0; i < len(parts)-1; i++ {
				curLine += parts[i]
				out <- curLine
				curLine = "" 	
			}

			curLine += parts[len(parts)-1]
		}

		if curLine != "" {
			out <- curLine
		}
	}()

	return out
}


func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

	for line := range getLinesChannel(f) {
		fmt.Printf("read: %s\n", line)
	}
}
