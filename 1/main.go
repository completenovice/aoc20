package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var numbers []int

func main() {
	var err error
	file, err := os.Open("input1.txt")
	defer func() {
		file.Close()
	}()

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}

	satisfied := false
	for _, one := range numbers {
		for _, two := range numbers {
			for _, three := range numbers {
				if one+two+three == 2020 {
					fmt.Println(one)
					fmt.Println(two)
					fmt.Println(three)
					fmt.Println(one * two * three)
					satisfied = true
					break
				}
			}
			if satisfied {
				break
			}
		}
		if satisfied {
			break
		}
	}

}
