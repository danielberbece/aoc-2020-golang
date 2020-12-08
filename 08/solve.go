package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	op     string
	p1     int
	wasRan bool
}

func runCodeUntilLoop(code []instruction) (int, bool) {
	accValue := 0
	wasForceClosed := false

	for ip := 0; ip < len(code); {
		if code[ip].wasRan {
			wasForceClosed = true
			break
		}
		code[ip].wasRan = true
		switch code[ip].op {
		case "acc":
			accValue += code[ip].p1
			ip++
		case "jmp":
			ip += code[ip].p1
		default:
			ip++
		}
	}

	return accValue, wasForceClosed
}

func switchOp(code []instruction, ip int) {
	if code[ip].op == "jmp" {
		code[ip].op = "nop"
	} else if code[ip].op == "nop" {
		code[ip].op = "jmp"
	}
}

func resetCode(code []instruction) {
	for i := 0; i < len(code); i++ {
		code[i].wasRan = false
	}
}

func readCode(filepath string) []instruction {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	code := make([]instruction, 0)

	for scanner.Scan() {
		x := scanner.Text()
		tokens := strings.Split(x, " ")
		n, _ := strconv.Atoi(tokens[1])
		code = append(code, instruction{tokens[0], n, false})
	}

	return code
}

func part1(code []instruction) {
	accValue, _ := runCodeUntilLoop(code)
	resetCode(code)
	fmt.Printf("Part 1: %d\n", accValue)
}

func part2(code []instruction) {
	accValue := 0
	wasForceClosed := false
	for i := 0; i < len(code); i++ {
		if code[i].op == "jmp" || code[i].op == "nop" {
			switchOp(code, i)
			accValue, wasForceClosed = runCodeUntilLoop(code)
			resetCode(code)
			if !wasForceClosed {
				break
			}
			switchOp(code, i)
		}
	}
	fmt.Printf("Part 2: %d\n", accValue)
}

func main() {
	code := readCode("./input")

	part1(code)
	part2(code)
}
