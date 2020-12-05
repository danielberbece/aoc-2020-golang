package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxID := 0
	minID := 127*8 + 7

	seats := make(map[int]bool)

	for scanner.Scan() {
		x := scanner.Text()
		left := 0
		right := 127
		for i := 0; i < 7; i++ {
			if x[i] == 'F' {
				right = (left + right) / 2
			} else {
				left = (left + right + 1) / 2
			}
		}
		id := 8 * right

		left = 0
		right = 7
		for i := 7; i < 10; i++ {
			if x[i] == 'L' {
				right = (left + right) / 2
			} else {
				left = (left + right + 1) / 2
			}
		}

		id += right
		seats[id] = true
		if id > maxID {
			maxID = id
		}
		if id < minID {
			minID = id
		}
	}

	missingID := -1
	for i := minID; i <= maxID; i++ {
		if _, ok := seats[i]; !ok {
			missingID = i
			break
		}
	}

	fmt.Printf("Part 1: %d\n", maxID)
	if missingID != -1 {
		fmt.Printf("Part 2: %d\n", missingID)
	} else {
		fmt.Println("Part 2: No missing ID found")
	}
}
