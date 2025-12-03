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
