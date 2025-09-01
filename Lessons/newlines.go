package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal(err)
	}

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
			fmt.Printf("read: %s\n", curLine)
			curLine = "" 	
		}

		curLine += parts[len(parts)-1]
	}

	if curLine != "" {
		fmt.Printf("read: %s\n", curLine)
	}
}
