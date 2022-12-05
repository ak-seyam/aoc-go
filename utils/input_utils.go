package utils

import (
	"bufio"
	"log"
	"os"
)

func GetInput(path string, onNewLine func(inp string)) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		onNewLine(line)
	}
}
