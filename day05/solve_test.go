package day05_test

import (
	"burlea/Advent_of_Code_2025/day05"
	"burlea/Advent_of_Code_2025/utils"
	"testing"
)

func TestSolvePuzzle1(t *testing.T) {
	input := utils.ReadInput()
	result := day05.SolvePuzzle1(input)
	t.Log(result)
}

func TestSolvePuzzle2(t *testing.T) {
	input := utils.ReadInput()
	result := day05.SolvePuzzle2(input)
	t.Log(result)
}
