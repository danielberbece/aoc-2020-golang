package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func createPartialSumsArray(arr []int) []int {
	partialSumsArray := make([]int, 0)
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		partialSumsArray = append(partialSumsArray, sum)
	}

	return partialSumsArray
}

func findContigousSectionSum(partialSumsArray []int, targetSum int) (int, int) {
	sumsDex := make(map[int]int)
	for i := 0; i < len(partialSumsArray); i++ {
		if pos, ok := sumsDex[partialSumsArray[i]-targetSum]; ok {
			return pos + 1, i
		}
		sumsDex[partialSumsArray[i]] = i
	}

	return -1, -1
}

func getMinAndMaxValues(arr []int) (int, int) {
	var min int
	var max int
	for i, elem := range arr {
		if i == 0 || elem < min {
			min = elem
		}
		if i == 0 || elem > max {
			max = elem
		}
	}

	return min, max
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbersDex := make(map[int]int)
	numbersArray := make([]int, 0)
	preambleLength := 25
	var vulnNum int

	i := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if i >= preambleLength {
			canBeSummedWithBeforeValues := false
			for j := i - 1; j >= i-preambleLength; j-- {
				if pos, ok := numbersDex[num-numbersArray[j]]; ok && pos >= i-preambleLength {
					canBeSummedWithBeforeValues = true
					break
				}
			}
			if !canBeSummedWithBeforeValues {
				vulnNum = num
				break
			}
		}
		numbersArray = append(numbersArray, num)
		numbersDex[num] = i
		i++
	}

	fmt.Printf("Part1: %v\n", vulnNum)

	partialSumsArray := createPartialSumsArray(numbersArray)
	leftIndex, rightIndex := findContigousSectionSum(partialSumsArray, vulnNum)

	if leftIndex == -1 {
		fmt.Println("Part2: no solution found!")
	} else {
		min, max := getMinAndMaxValues(numbersArray[leftIndex : rightIndex+1])
		fmt.Printf("Part 2: %d\n", min+max)
	}
}
