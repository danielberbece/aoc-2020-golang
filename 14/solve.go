package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const part1 = 1
const part2 = 2

func readInput(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func applyMask(num uint64, mask string) uint64 {
	var newNum uint64
	for i := 0; i < len(mask); i++ {
		if mask[len(mask)-i-1] == 'X' {
			newNum += num & (1 << i)
		} else if mask[len(mask)-i-1] == '1' {
			newNum += (1 << i)
		}
	}
	return newNum
}

func translateAddr(addr string) uint64 {
	var uintAddr uint64
	for i := 0; i < len(addr); i++ {
		if addr[len(addr)-i-1] == '1' {
			uintAddr += (1 << i)
		}
	}

	return uintAddr
}

func updateFloatingAddress(formedAddr string, floatingAddr string, num uint64, memory map[uint64]uint64) {
	pos := len(formedAddr)
	for pos < len(floatingAddr) && floatingAddr[pos] != 'X' {
		formedAddr += string(floatingAddr[pos])
		pos++
	}

	if pos == len(floatingAddr) {
		addr := translateAddr(formedAddr)
		memory[addr] = num
		return
	}

	updateFloatingAddress(formedAddr+"1", floatingAddr, num, memory)
	updateFloatingAddress(formedAddr+"0", floatingAddr, num, memory)
}

func applyAddressMask(baseAddr uint64, mask string) string {
	floatingAddr := ""

	for i := 0; i < len(mask); i++ {
		if mask[i] == '0' {
			floatingAddr = floatingAddr + strconv.FormatUint((baseAddr>>(len(mask)-i-1))%2, 10)
		} else if mask[i] == '1' {
			floatingAddr = floatingAddr + "1"
		} else {
			floatingAddr = floatingAddr + "X"
		}
	}

	return floatingAddr
}

func runInput(input []string, part int) map[uint64]uint64 {
	memory := make(map[uint64]uint64)
	var mask string
	for i := 0; i < len(input); i++ {
		if input[i][1] == 'a' { // mask
			mask = strings.Split(input[i], " ")[2]
		} else if input[i][1] == 'e' { // memory
			tokens := strings.Split(input[i], " ")
			addr, _ := strconv.ParseUint(tokens[0][4:len(tokens[0])-1], 10, 64)
			num, _ := strconv.ParseUint(tokens[2], 10, 64)
			if part == part1 {
				numWithMask := applyMask(num, mask)
				memory[addr] = numWithMask
			} else {
				floatingAddr := applyAddressMask(addr, mask)
				updateFloatingAddress("", floatingAddr, num, memory)
			}
		}
	}

	return memory
}

func getMemorySum(memory map[uint64]uint64) uint64 {
	var sum uint64
	for _, v := range memory {
		sum += v
	}

	return sum
}

func main() {
	input := readInput("./input")
	memory := runInput(input, part1)
	fmt.Printf("Part 1: %v\n", getMemorySum(memory))

	memory = runInput(input, part2)
	fmt.Printf("Part 2: %v\n", getMemorySum(memory))
}
