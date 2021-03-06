package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var lines []string

func main() {
	var err error
	file, err := os.Open("input.txt")
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

	validCount := 0
	for _, line := range lines {
		fmt.Printf("%s\n", line)
		min, max, ch, password := parser(line)
		if isValid(min, max, ch, password) {
			fmt.Println(line)
			validCount = validCount + 1
		}
	}
	fmt.Println(validCount)
}

func parser(rule string) (int, int, string, []string) {
	parts := strings.Split(rule, " ")

	numRange := parts[0]
	minS := strings.Split(numRange, "-")[0]
	min, _ := strconv.Atoi(minS)
	maxS := strings.Split(numRange, "-")[1]
	max, _ := strconv.Atoi(maxS)

	ch := parts[1]
	ch = strings.Split(ch, ":")[0]
	pw := parts[2]
	password := []string{}
	for _, c := range []rune(pw) {
		password = append(password, string(c))
	}

	return min, max, ch, password
}

func isValid(min, max int, ch string, password []string) bool {
	fmt.Printf("Checking '%s' in '%s' is in position %d and less than %d\n", ch, password, min, max)
	if ch == password[min-1] && ch == password[max-1] {
		return false
	}
	if ch == password[min-1] || ch == password[max-1] {
		fmt.Println("yep!")
		return true
	}
	return false
}
