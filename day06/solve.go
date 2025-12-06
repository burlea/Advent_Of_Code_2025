package day06

import (
	"burlea/Advent_of_Code_2025/utils"
)

func SolvePuzzle1(input []byte) uint64 {
	lines := utils.ParseInputTo2DStringSpacedArray(input)
	formattedLines := utils.ParseSpacesFromLines(lines)

	var totalResult uint64 = 0
	for j := range formattedLines[0] {

		var numbers []uint64
		for i := 0; i < len(formattedLines)-1; i++ {
			numbers = append(numbers, utils.ConvertStringToUInt64(formattedLines[i][j]))
		}

		operand := formattedLines[len(formattedLines)-1][j]

		var columnTotal uint64
		if operand == "*" {
			columnTotal = 1
		}

		for _, number := range numbers {
			switch operand {
			case "+":
				columnTotal += number
			case "*":
				columnTotal *= number
			}
		}
		totalResult += columnTotal
	}
	return totalResult
}

func SolvePuzzle2(input []byte) uint64 {
	lines := utils.ParseInputTo2DStringArray(input)

	operandLine := lines[len(lines)-1]

	var operandIndicies []int

	for index := range operandLine {
		if operandLine[index] == "+" || operandLine[index] == "*" {
			operandIndicies = append(operandIndicies, index)
		}
	}

	operandIndicies = append(operandIndicies, len(operandLine))

	var totalResult uint64

	for i := 0; i < len(operandIndicies)-1; i++ {
		startOperandIndex := operandIndicies[i]
		endOperandIndex := operandIndicies[i+1]
		operand := operandLine[startOperandIndex]

		var numbers []uint64

		for col := startOperandIndex; col < endOperandIndex; col++ {

			var numberString string

			for row := 0; row < len(lines)-1; row++ {
				if lines[row][col] != "" && lines[row][col] != " " {
					numberString += lines[row][col]
				}
			}

			numbers = append(numbers, utils.ConvertStringToUInt64(numberString))
		}

		if i != len(operandIndicies)-2 {
			numbers = numbers[:len(numbers)-1]
		}

		var operandTotal uint64

		if operand == "*" {
			operandTotal = 1
		}

		for _, number := range numbers {
			switch operand {
			case "+":
				operandTotal += number
			case "*":
				operandTotal *= number
			}
		}
		totalResult += operandTotal
	}
	return totalResult
}
