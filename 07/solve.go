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

type innerBag struct {
	color string
	num   int
}

type bag struct {
	childBags       []innerBag
	parentBagColors []string
}

var bagsDex = make(map[string]bag)

func createBag() bag {
	return bag{nil, nil}
}

func removeLastWord(str string) string {
	tokens := strings.Split(str, " ")
	tokens = tokens[:len(tokens)-1]
	return strings.Join(tokens, " ")
}

func removeFirstWord(str string) string {
	tokens := strings.Split(str, " ")
	tokens = tokens[1:len(tokens)]
	return strings.Join(tokens, " ")
}

func getFirstNumber(str string) (int, error) {
	tokens := strings.Split(str, " ")
	return strconv.Atoi(tokens[0])
}

func addBagToDex(color string) bag {
	if parentBag, ok := bagsDex[color]; !ok {
		parentBag = createBag()
		bagsDex[color] = parentBag
	}

	return bagsDex[color]
}

func addChildToBag(parentColor string, childColor string, numBags int) {
	parentBag := bagsDex[parentColor]
	newChildBag := innerBag{childColor, numBags}
	parentBag.childBags = append(parentBag.childBags, newChildBag)
	bagsDex[parentColor] = parentBag
}

func parseTokens(tokens []string) {
	parentBagColor := removeLastWord(tokens[0])
	addBagToDex(parentBagColor)

	for i := 1; i < len(tokens); i++ {
		tokens[i] = removeLastWord(tokens[i])
		if tokens[i] != "no other" {
			numBags, _ := getFirstNumber(tokens[i])
			childBagColor := removeFirstWord(tokens[i])
			addChildToBag(parentBagColor, childBagColor, numBags)
			childBag := addBagToDex(childBagColor)
			childBag.parentBagColors = append(childBag.parentBagColors, parentBagColor)
			bagsDex[childBagColor] = childBag
		}
	}
}

var parentBagsDex = make(map[string]bool)

func countParentsOf(color string) int {
	parentBags := bagsDex[color].parentBagColors

	total := 0
	for i := 0; i < len(parentBags); i++ {
		if _, ok := parentBagsDex[parentBags[i]]; !ok {
			parentBagsDex[parentBags[i]] = true
			total += 1 + countParentsOf(parentBags[i])
		}
	}

	return total
}

func countChildrenOf(color string) int {
	childBags := bagsDex[color].childBags

	total := 0
	for i := 0; i < len(childBags); i++ {
		total += childBags[i].num + childBags[i].num*countChildrenOf(childBags[i].color)
	}

	return total
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(` contain |, |\.`)

	for scanner.Scan() {
		line := scanner.Text()
		tokens := re.Split(line, -1)
		tokens = tokens[:len(tokens)-1]
		parseTokens(tokens)
	}

	numParents := countParentsOf("shiny gold")
	fmt.Printf("Part 1: %d\n", numParents)
	numChilds := countChildrenOf("shiny gold")
	fmt.Printf("Part 2: %d\n", numChilds)
}
