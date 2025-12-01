package day01

import (
	"burlea/Advent_of_Code_2025/utils"
	"math"
	"strconv"
)

func SolvePuzzle1(input []byte) int {
	// TODO: solve puzzle 1
	rotations := utils.ParseInputToStringArray(input)

	currentNumber := 50
	totalZeros := 0

	for _, rotation := range rotations {
		value, _ := strconv.Atoi(rotation[1:])
		switch rotation[0] {
		case 'R':
			currentNumber += value
		case 'L':
			currentNumber -= value
		}
		if currentNumber%100 == 0 {
			totalZeros++
		}
	}
	return totalZeros
}

func SolvePuzzle2(input []byte) int {
	// TODO: solve puzzle 2
	rotations := utils.ParseInputToStringArray(input)

	var currentNumber int = 50
	totalZeros := 0

	for _, rotation := range rotations {
		value, _ := strconv.Atoi(rotation[1:])
		var prevValue int = currentNumber

		switch rotation[0] {
		case 'R':
			currentNumber += value
		case 'L':
			currentNumber -= value
		}

		if prevValue/100 != currentNumber/100 {
			totalZeros += int(math.Abs(float64(prevValue/100) - float64(currentNumber/100)))
		}
		if (currentNumber < 0 && prevValue > 0) || (currentNumber > 0 && prevValue < 0 || (currentNumber == 0 && prevValue != 0)) {
			totalZeros++
		}

		currentNumber = currentNumber % 100
	}
	return totalZeros
}
