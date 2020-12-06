package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countIfBiggerThan(array []int, lowerBound int) int {
	result := 0
	for _, v := range array {
		if v > lowerBound {
			result++
		}
	}
	return result
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	questionsArray := make([]int, 26)
	totalQuestionsAnswered := 0
	totalQuestionsAnsweredByAll := 0
	peopleInGroup := 0

	for scanner.Scan() {
		x := scanner.Text()
		if x == "" {
			questionsAnswered := countIfBiggerThan(questionsArray, 0)
			totalQuestionsAnswered += questionsAnswered
			questionsAnswered = countIfBiggerThan(questionsArray, peopleInGroup-1)
			totalQuestionsAnsweredByAll += questionsAnswered

			peopleInGroup = 0
			questionsArray = make([]int, 26)
		} else {
			peopleInGroup++
			for _, c := range x {
				questionsArray[c-'a']++
			}
		}
	}

	questionsAnswered := countIfBiggerThan(questionsArray, 0)
	totalQuestionsAnswered += questionsAnswered
	questionsAnswered = countIfBiggerThan(questionsArray, peopleInGroup-1)
	totalQuestionsAnsweredByAll += questionsAnswered

	fmt.Printf("Part 1: %d\n", totalQuestionsAnswered)
	fmt.Printf("Part 2: %d\n", totalQuestionsAnsweredByAll)
}
