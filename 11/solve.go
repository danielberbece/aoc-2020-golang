package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
)

const part1 = 1
const part2 = 2

func readSeats(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	seats := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		seats = append(seats, line)
	}

	return seats
}

var adjacentDirections = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func getOccupiedAdjacent(seats []string, row int, col int, part int) int {
	numOccupiedSeats := 0
	for _, direction := range adjacentDirections {
		pos := make([]int, 2)
		for {
			pos[0] += direction[0]
			pos[1] += direction[1]
			// If inside the area of the seats
			if row+pos[0] >= 0 && row+pos[0] < len(seats) && col+pos[1] >= 0 && col+pos[1] < len(seats[row]) {
				if seats[row+pos[0]][col+pos[1]] == '#' {
					numOccupiedSeats++
					break

					// On part 2 look only at the first sight seat in a direction
				} else if seats[row+pos[0]][col+pos[1]] == 'L' && part == part2 {
					break
				}
			} else { // If reached the outside of the seats space
				break
			}

			// On part 1 look only at the first neighbour in a direction
			if part == part1 {
				break
			}
		}

	}

	return numOccupiedSeats
}

func applyRules(seats []string, part int) []string {
	newSeats := make([]string, len(seats))
	copy(newSeats, seats)

	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '.' {
				continue
			}

			occupiedAdjacent := getOccupiedAdjacent(seats, i, j, part)
			if seats[i][j] == 'L' && occupiedAdjacent == 0 {
				newSeats[i] = newSeats[i][:j] + "#" + newSeats[i][j+1:]
			} else if (part == part1 && seats[i][j] == '#' && occupiedAdjacent >= 4) ||
				(part == part2 && seats[i][j] == '#' && occupiedAdjacent >= 5) {
				newSeats[i] = newSeats[i][:j] + "L" + newSeats[i][j+1:]
			}
		}
	}

	return newSeats
}

func simulateSeating(seats []string, part int) []string {
	newSeats := applyRules(seats, part)

	if reflect.DeepEqual(newSeats, seats) {
		return seats
	}

	return simulateSeating(newSeats, part)
}

func countAllOccupiedSeats(seats []string) int {
	occupiedSeats := 0
	for i := 0; i < len(seats); i++ {
		for j := 0; j < len(seats[i]); j++ {
			if seats[i][j] == '#' {
				occupiedSeats++
			}
		}
	}

	return occupiedSeats
}

func main() {
	seats := readSeats("./input")

	finalSeats := simulateSeating(seats, part1)
	fmt.Printf("Part 1: %v\n", countAllOccupiedSeats(finalSeats))

	finalSeats = simulateSeating(seats, part2)
	fmt.Printf("Part 2: %v\n", countAllOccupiedSeats(finalSeats))
}
