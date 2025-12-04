package day03

import (
	"burlea/Advent_of_Code_2025/utils"
	"strconv"
	"strings"
)

func SolvePuzzle1(input []byte) int {
	banks := utils.ParseInputTo2DStringArray(input)

	var joltageTotal int

	for _, bank := range banks {

		var highestNumber int = 0
		var secondHighestNumber int = 0

		for i := 0; i < len(bank)-1; i++ {
			joltage, _ := strconv.Atoi(bank[i])
			nextJoltage, _ := strconv.Atoi(bank[i+1])

			if joltage > highestNumber {
				highestNumber = joltage
				secondHighestNumber = nextJoltage
			} else if joltage > secondHighestNumber {
				secondHighestNumber = joltage
			}
		}

		lastJoltage, _ := strconv.Atoi(bank[len(bank)-1])
		if lastJoltage > secondHighestNumber {
			secondHighestNumber = lastJoltage
		}

		joltageTotal += highestNumber*10 + secondHighestNumber
	}
	return joltageTotal
}

func SolvePuzzle2(input []byte) uint64 {
	banks := utils.ParseInputTo2DStringArray(input)

	var joltageTotal uint64

	for _, bank := range banks {
		const target = 12
		n := len(bank)

		if n <= target {
			num := make([]string, target)
			copy(num, bank)
			for i := len(bank); i < target; i++ {
				num[i] = "0"
			}
			joltageTotal += utils.ConvertStringToUInt64(strings.Join(num, ""))
			continue
		}

		toRemove := n - target
		stack := make([]string, 0, n)

		for i := 0; i < n; i++ {
			curr := bank[i]
			for toRemove > 0 && len(stack) > 0 && stack[len(stack)-1] < curr {
				stack = stack[:len(stack)-1]
				toRemove--
			}
			stack = append(stack, curr)
		}

		if toRemove > 0 {
			stack = stack[:len(stack)-toRemove]
		}

		result := stack[:target]
		joltageTotal += utils.ConvertStringToUInt64(strings.Join(result, ""))
	}
	return joltageTotal
}
