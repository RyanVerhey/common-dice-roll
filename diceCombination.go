package main

import (
	"sort"
)

type diceCombination []int

func generateDiceCombinations(numOfDice int, numOfSides int) []diceCombination {
	if numOfDice <= 0 {
		return []diceCombination{}
	}
	var rolls []diceCombination

	if numOfDice == 1 {
		for i := 1; i <= numOfSides; i++ {
			rolls = append(rolls, diceCombination{i})
		}

		return rolls
	}

	previousRolls := generateDiceCombinations(numOfDice-1, numOfSides)

	for _, roll := range previousRolls {
		for i := 1; i <= numOfSides; i++ {
			newRoll := append(roll, i)
			rolls = append(rolls, newRoll)
		}
	}

	return rolls
}

func sumDiceRolls(combinations []diceCombination) []int {
	sums := make([]int, len(combinations))
	for i, numArr := range combinations {
		sums[i] = sum(numArr)
	}

	return sums
}

func getDiceRollsWithLargestProbability(diceRolls []diceCombination) []int {
	diceRollSums := sumDiceRolls(diceRolls)
	rollsGroupedByCount := groupByCount(diceRollSums)

	keys := make([]int, 0, len(rollsGroupedByCount))
	for k, _ := range rollsGroupedByCount {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	largestCount := keys[len(keys)-1]

	return rollsGroupedByCount[largestCount]
}

func sum(arr []int) int {
	result := 0

	for _, v := range arr {
		result += v
	}

	return result
}

func groupByCount(s []int) map[int]diceCombination {
	counts := make(map[int]int)
	for _, elem := range s {
		counts[elem] += 1
	}

	rollsByCount := make(map[int]diceCombination)
	for sum, count := range counts {
		rollsByCount[count] = append(rollsByCount[count], sum)
	}

	return rollsByCount
}
