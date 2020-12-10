package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func readAdapters(filepath string) []int {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	adapters := make([]int, 0)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		adapters = append(adapters, n)
	}

	return adapters
}

func getAdapterDifferences(adapters []int) []int {
	diffs := []int{0, 0, 0, 0}

	// charging outlet to adapter difference
	diffs[adapters[0]-0]++
	// between adapters difference
	for i := 0; i < len(adapters)-1; i++ {
		diffs[adapters[i+1]-adapters[i]]++
	}
	// adapter to device difference
	diffs[3]++

	return diffs
}

func getNumArrangements(adapters []int) int {
	adaptersExtended := make([]int, 1)
	adaptersExtended = append(adaptersExtended, adapters...)
	adaptersExtended = append(adaptersExtended, adapters[len(adapters)-1]+3)
	d := make([]int, len(adaptersExtended))
	d[0] = 1

	for i := 1; i < len(adaptersExtended); i++ {
		for j := i - 1; j >= 0 && adaptersExtended[j] >= adaptersExtended[i]-3; j-- {
			d[i] += d[j]
		}
	}

	return d[len(d)-1]
}

func main() {
	adapters := readAdapters("./input")
	sort.Ints(adapters)

	diffs := getAdapterDifferences(adapters)
	fmt.Printf("Part 1: %v\n", diffs[1]*diffs[3])

	arrangements := getNumArrangements(adapters)
	fmt.Printf("Part 2: %v\n", arrangements)
}
