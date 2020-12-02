package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func countLetterInString(str string, letter string) int {
	n := 0
	for i := 0; i < len(str); i++ {
		if str[i] == letter[0] {
			n++
		}
	}
	return n
}

func checkPositions(str string, letter string, pos1 int, pos2 int) bool {
	if (str[pos1] == letter[0] && str[pos2] != letter[0]) || (str[pos1] != letter[0] && str[pos2] == letter[0]) {
		return true
	}

	return false
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	zp := regexp.MustCompile(`(-)|( )|(: )`)
	cntPart1 := 0
	cntPart2 := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		x := scanner.Text()
		if err != nil {
			break
		}
		tokens := zp.Split(x, -1)
		minLetters, _ := strconv.Atoi(tokens[0])
		maxLetters, _ := strconv.Atoi(tokens[1])
		letterCnt := countLetterInString(tokens[3], tokens[2])

		if letterCnt >= minLetters && letterCnt <= maxLetters {
			cntPart1++
		}

		if checkPositions(tokens[3], tokens[2], minLetters-1, maxLetters-1) {
			cntPart2++
		}
	}

	fmt.Printf("Part 1: %d\n", cntPart1)
	fmt.Printf("Part 2: %d\n", cntPart2)
}
