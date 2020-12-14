package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func readInput(filepath string) (int, []int) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timestamp, _ := strconv.Atoi(scanner.Text())
	buses := make([]int, 0)

	re := regexp.MustCompile(",")
	for scanner.Scan() {
		line := re.Split(scanner.Text(), -1)
		for _, bus := range line {
			if bus != "x" {
				busID, _ := strconv.Atoi(bus)
				buses = append(buses, busID)
			}
		}
	}

	return timestamp, buses
}

func main() {
	timestamp, buses := readInput("./input")
	minBus := []int{buses[0], buses[0] - (timestamp % buses[0])}
	for _, bus := range buses {
		currentBusDiff := bus - (timestamp % bus)
		if minBus[1] > currentBusDiff {
			minBus = []int{bus, currentBusDiff}
		}
	}
	fmt.Printf("Part 1: %v\n", minBus[0]*minBus[1])

}
