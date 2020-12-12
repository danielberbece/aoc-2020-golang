package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

const part1 = 1
const part2 = 2

func readNavInstructions(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	navInstructions := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		navInstructions = append(navInstructions, line)
	}

	return navInstructions
}

func rotateWaypoint(position [3]int, degrees int) [3]int {
	waypointPosition := [2]int{position[0], position[1]}
	newWaypointPosition := [2]int{0, 0}

	sin := int(math.Sin(math.Pi * float64(degrees) / 180))
	cos := int(math.Cos(math.Pi * float64(degrees) / 180))

	newWaypointPosition[0] = waypointPosition[0]*cos - waypointPosition[1]*sin
	newWaypointPosition[1] = waypointPosition[0]*sin + waypointPosition[1]*cos

	position[0] = newWaypointPosition[0]
	position[1] = newWaypointPosition[1]
	return position
}

func navigate(navInstructions []string, part int) [2]int {
	var position [3]int = [3]int{0, 0, 0}
	var shipPosition [2]int = [2]int{0, 0}
	if part == part2 {
		position = [3]int{10, 1, 0}
	}

	for i := 0; i < len(navInstructions); i++ {
		action := navInstructions[i][0]
		value, _ := strconv.Atoi(navInstructions[i][1:])
		switch action {
		case 'E':
			position[0] += value
		case 'W':
			position[0] -= value
		case 'N':
			position[1] += value
		case 'S':
			position[1] -= value
		case 'L':
			if part == part2 {
				position = rotateWaypoint(position, value)
			} else {
				position[2] += value
				position[2] %= 360
			}
		case 'R':
			if part == part2 {
				position = rotateWaypoint(position, -value)
			} else {
				position[2] -= value

				for position[2] < 0 {
					position[2] = 360 + position[2]
				}
			}
		case 'F':
			if part == part1 {
				switch position[2] {
				case 0:
					position[0] += value
				case 90:
					position[1] += value
				case 180:
					position[0] -= value
				case 270:
					position[1] -= value
				}
			} else {
				shipPosition[0] += value * position[0]
				shipPosition[1] += value * position[1]
			}
		}
	}

	if part == part1 {
		return [2]int{position[0], position[1]}
	}

	return shipPosition
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	navInstructions := readNavInstructions("./input")

	finalPosition := navigate(navInstructions, part1)
	fmt.Printf("Part 1: %v\n", abs(finalPosition[0])+abs(finalPosition[1]))

	finalPosition2 := navigate(navInstructions, part2)
	fmt.Printf("Part 2: %v\n", abs(finalPosition2[0])+abs(finalPosition2[1]))
}
