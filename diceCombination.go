package main

import (
	"math"
	"sort"
)

type diceCombination []int

func generateDiceCombinations(numOfDice int, numOfSides int) []diceCombination {
	numOfCombinations := int(math.Pow(float64(numOfSides), float64(numOfDice)))
	combinations := make([]diceCombination, numOfCombinations, numOfCombinations)

	for i := 0; i < numOfSides; i++ {
		combinations[i] = []int{(i + 1)}
	}

	for diceI := 1; diceI < numOfDice; diceI++ {
		holding := make([]diceCombination, 0)

		for sideI := 0; sideI < numOfSides; sideI++ {
			comboCopy := make([]diceCombination, len(combinations))
			copy(comboCopy, combinations)
			for index, elem := range comboCopy {
				comboCopy[index] = append([]int{(sideI + 1)}, elem...)
			}
			holding = append(holding, comboCopy...)
		}
		combinations = holding
	}

	return combinations
}

func getDiceRollsWithLargestProbability(combinations []diceCombination) []int {
	sums := make([]int, len(combinations))
	for i, numArr := range combinations {
		sums[i] = sum(numArr)
	}

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
	return sumsByCount[largestCount]
}

func sum(arr []int) int {
	result := 0

	for _, v := range arr {
		result += v
	}

	return result
}
