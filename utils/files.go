package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func ReadInput() []byte {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("Failed to get caller information")
	}

	filePath := fmt.Sprintf("%s/input.txt", filepath.Dir(file))
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("File not found for path: %s", err.Error()))
	}

	return content
}

func ParseInputToString(input []byte) string {
	return string(input)
}

func ParseInputToStringArray(input []byte) []string {
	return strings.Split(string(input), "\n")
}

func ParseInputCommaSeparatedArray(input []byte) []string {
	return strings.Split(string(input), ",")
}

func ParseInputTo2DStringArray(input []byte) [][]string {
	rows := ParseInputToStringArray(input)
	var result [][]string
	for _, row := range rows {
		cols := strings.Split(row, "")

		result = append(result, cols)
	}
	return result
}

func ParseInputTo2DStringSpacedArray(input []byte) [][]string {
	rows := ParseInputToStringArray(input)
	var result [][]string
	for _, row := range rows {
		cols := strings.Split(row, " ")

		result = append(result, cols)
	}
	return result
}

func ConvertStringToUInt64(str string) uint64 {
	num, _ := strconv.ParseUint(str, 10, 64)
	return num
}

func ConvertStringToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func ConvertIntToString(num int) string {
	return strconv.Itoa(num)
}

func IsWithinBounds(i int, j int, totalRows int, totalCols int) bool {
	return i >= 0 && j >= 0 && i < totalRows && j < totalCols
}

func ParseToTwoPartsSpaced(input []byte) ([]string, []string) {
	database := ParseInputToStringArray(input)

	var part1 []string
	var part2 []string

	var indexToStartPart2 int

	for i, line := range database {
		if line == "" {
			indexToStartPart2 = i + 1
			break
		}

		part1 = append(part1, line)
	}

	for i := indexToStartPart2; i < len(database); i++ {
		line := database[i]
		part2 = append(part2, line)
	}

	return part1, part2
}
