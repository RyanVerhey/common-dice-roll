package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	args := os.Args[1:]

	pattern := regexp.MustCompile("([0-9]+)d([0-9]+)")
	match := pattern.FindStringSubmatch(args[0])

	if len(match) == 0 {
		log.Fatal("Input must match NdN")
	}

	diceNum, _ := strconv.Atoi(match[1])
	sideNum, _ := strconv.Atoi(match[2])

	combinations := generateDiceCombinations(diceNum, sideNum)

	largestProbabilities := getDiceRollsWithLargestProbability(combinations)

	fmt.Printf("Dice rolls with the largest probability: %d", largestProbabilities)
}
