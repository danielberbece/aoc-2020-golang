package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getTreesMap() []string {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var treesMap []string
	for scanner.Scan() {
		x := scanner.Text()
		treesMap = append(treesMap, x)
	}

	return treesMap
}

func getNumTreesOnSlope(treesMap []string, stepX int, stepY int) int {
	mapWidth := len(treesMap[0])
	mapHeight := len(treesMap)

	posX := 0
	posY := 0

	numTreesTouched := 0

	for posY < mapHeight {
		if treesMap[posY][posX] == '#' {
			numTreesTouched++
		}
		posY += stepY
		posX += stepX
		posX %= mapWidth
	}

	return numTreesTouched
}

func main() {
	treesMap := getTreesMap()

	// Part 1
	treesSlope31 := getNumTreesOnSlope(treesMap, 3, 1)
	fmt.Printf("Part 1: %d\n", treesSlope31)

	// Part 2
	treesSlope11 := getNumTreesOnSlope(treesMap, 1, 1)
	treesSlope51 := getNumTreesOnSlope(treesMap, 5, 1)
	treesSlope71 := getNumTreesOnSlope(treesMap, 7, 1)
	treesSlope12 := getNumTreesOnSlope(treesMap, 1, 2)

	probability := treesSlope11 * treesSlope12 * treesSlope31 * treesSlope51 * treesSlope71
	fmt.Printf("Part 2: %d\n", probability)
}
