package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
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
	numOfCombinations := int(math.Pow(float64(sideNum), float64(diceNum)))

	// fmt.Println(diceNum, sideNum, numOfCombinations)

	combinations := make([][]int, 0, numOfCombinations)
	diceI := 0
	sideI := 0
	// setting up combinations
	for sideI < sideNum {
		combinations = append(combinations, []int{(sideI + 1)})

		sideI++
	}
	sideI = 0

	for diceI < (diceNum - 1) {
		comboHolding := make([][]int, 0)
		for sideI < sideNum {
			comboCopy := make([][]int, len(combinations))
			copy(comboCopy, combinations)
			for index, elem := range comboCopy {
				comboCopy[index] = append([]int{(sideI + 1)}, elem...)
			}
			comboHolding = append(comboHolding, comboCopy...)
			sideI++
		}
		sideI = 0
		combinations = comboHolding

		diceI++
	}

	// fmt.Println(combinations, len(combinations))

	sums := make([]int, numOfCombinations)
	for i, numArr := range combinations {
		sums[i] = sum(numArr)
	}
	// fmt.Println(sums, len(sums))
	counts := make(map[int]int)
	for _, sum := range sums {
		counts[sum] += 1
	}
	// fmt.Println(counts)
	sumsByCount := make(map[int][]int)
	for sum, count := range counts {
		sumsByCount[count] = append(sumsByCount[count], sum)
	}

	keys := make([]int, 0, len(sumsByCount))
	for k, _ := range sumsByCount {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	largestCount := keys[len(keys)-1]
	fmt.Printf("Dice rolls with the largest probability: %d", sumsByCount[largestCount])
}

func sum(arr []int) int {
	result := 0

	for _, v := range arr {
		result += v
	}

	return result
}
