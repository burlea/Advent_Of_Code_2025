package day07

import (
	"burlea/Advent_of_Code_2025/utils"
	"fmt"
	"slices"
)

func SolvePuzzle1(input []byte) int {
	diagram := utils.ParseInputTo2DStringArray(input)

	currentLineIndexes := make([]int, len(diagram[0]))

	var currentSplits int

	for index, element := range diagram[0] {
		if element == "S" {
			currentLineIndexes = append(currentLineIndexes, index)
		}
	}

	for _, row := range diagram[1:] {

		nextLineIndexes := make([]int, len(diagram[0]))

		for _, currentIndexToCheck := range currentLineIndexes {
			if row[currentIndexToCheck] == "^" {
				leftSplit := currentIndexToCheck - 1
				rightSplit := currentIndexToCheck + 1

				if leftSplit >= 0 {
					if !slices.Contains(nextLineIndexes, leftSplit) {
						nextLineIndexes = append(nextLineIndexes, leftSplit)
					}
				}

				if rightSplit < len(diagram[0]) {
					if !slices.Contains(nextLineIndexes, rightSplit) {
						nextLineIndexes = append(nextLineIndexes, rightSplit)
					}
				}

				currentSplits++
			} else {
				if !slices.Contains(nextLineIndexes, currentIndexToCheck) {
					nextLineIndexes = append(nextLineIndexes, currentIndexToCheck)
				}
			}

		}
		currentLineIndexes = nextLineIndexes
	}

	fmt.Println(currentSplits)
	return currentSplits
}

func SolvePuzzle2(input []byte) int {
	diagram := utils.ParseInputTo2DStringArray(input)

	currentLineIndexes := make(map[int]int)

	for index, element := range diagram[0] {
		if element == "S" {
			currentLineIndexes[index] = 1
		}
	}

	for _, row := range diagram[1:] {

		nextLineIndexes := make(map[int]int)

		for currentIndexToCheck, currentPaths := range currentLineIndexes {
			if row[currentIndexToCheck] == "^" {
				leftSplit := currentIndexToCheck - 1
				rightSplit := currentIndexToCheck + 1

				if leftSplit >= 0 {
					val, contains := nextLineIndexes[leftSplit]
					if contains {
						nextLineIndexes[leftSplit] = val + currentPaths
					} else {
						nextLineIndexes[leftSplit] = currentPaths
					}
				}

				if rightSplit < len(diagram[0]) {
					val, contains := nextLineIndexes[rightSplit]
					if contains {
						nextLineIndexes[rightSplit] = val + currentPaths
					} else {
						nextLineIndexes[rightSplit] = currentPaths
					}
				}
			} else {
				val, contains := nextLineIndexes[currentIndexToCheck]
				if contains {
					nextLineIndexes[currentIndexToCheck] = val + currentPaths
				} else {
					nextLineIndexes[currentIndexToCheck] = currentPaths
				}
			}
		}
		currentLineIndexes = nextLineIndexes
	}

	var totalPaths int
	for _, value := range currentLineIndexes {
		totalPaths += value
	}

	return totalPaths
}
