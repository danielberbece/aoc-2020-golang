package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var eyeColors = map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}

func arePassportFieldsPresent(passport map[string]string) bool {
	for _, field := range requiredFields {
		if _, ok := passport[field]; !ok {
			return false
		}
	}
	return true
}

func isPassportValid(passport map[string]string) bool {
	if !arePassportFieldsPresent(passport) {
		return false
	}

	x, err := strconv.Atoi(passport["byr"])
	if err != nil || x < 1920 || x > 2002 {
		return false
	}

	x, err = strconv.Atoi(passport["iyr"])
	if err != nil || x < 2010 || x > 2020 {
		return false
	}

	x, err = strconv.Atoi(passport["eyr"])
	if err != nil || x < 2020 || x > 2030 {
		return false
	}

	if strings.HasSuffix(passport["hgt"], "in") {
		x, err = strconv.Atoi(strings.TrimSuffix(passport["hgt"], "in"))
		if err != nil || x < 59 || x > 76 {
			return false
		}
	} else if strings.HasSuffix(passport["hgt"], "cm") {
		x, err = strconv.Atoi(strings.TrimSuffix(passport["hgt"], "cm"))
		if err != nil || x < 150 || x > 193 {
			return false
		}
	} else {
		return false
	}

	match, err := regexp.MatchString(`^#[0-9a-f]{6}$`, passport["hcl"])
	if err != nil || !match {
		return false
	}

	if _, ok := eyeColors[passport["ecl"]]; !ok {
		return false
	}

	match, err = regexp.MatchString(`^[0-9]{9}$`, passport["pid"])
	if err != nil || !match {
		return false
	}

	return true
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numHasFieldsPassports := 0
	numValidPassports := 0
	currentPassport := make(map[string]string)
	zp := regexp.MustCompile(`:| `)

	for scanner.Scan() {
		x := scanner.Text()
		if x == "" {
			if arePassportFieldsPresent(currentPassport) {
				numHasFieldsPassports++
			}
			if isPassportValid(currentPassport) {
				numValidPassports++
			}
			currentPassport = make(map[string]string)
		} else {
			tokens := zp.Split(x, -1)
			for i := 0; i < len(tokens); i += 2 {
				currentPassport[tokens[i]] = tokens[i+1]
			}
		}
	}

	if arePassportFieldsPresent(currentPassport) {
		numHasFieldsPassports++
	}

	if isPassportValid(currentPassport) {
		numValidPassports++
	}

	fmt.Printf("Part 1: %d\n", numHasFieldsPassports)
	fmt.Printf("Part 2: %d\n", numValidPassports)
}
