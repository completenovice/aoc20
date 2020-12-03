package utils

import (
	"bufio"
	"log"
	"os"
)

var lines []string

// ReadLines returns a list of strings found in a file
func ReadLines(filepath string) ([]string, error) {
	var err error
	file, err := os.Open(filepath)
	defer func() {
		file.Close()
	}()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if err != nil {
			panic(err)
		}
		lines = append(lines, scanner.Text())
	}

	return lines, err

}
