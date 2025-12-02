package day04_test

import (
	"burlea/Advent_of_Code_2025/day04"
	"burlea/Advent_of_Code_2025/utils"
	"testing"
	"time"
)

func TestSolvePuzzle1(t *testing.T) {
	start := time.Now()
	input := utils.ReadInput()
	result := day04.SolvePuzzle1(input)
	t.Log(result)
	t.Log("Duration:", time.Since(start))
}

func TestSolvePuzzle2(t *testing.T) {
	start := time.Now()
	input := utils.ReadInput()
	result := day04.SolvePuzzle2(input)
	t.Log(result)
	t.Log("Duration:", time.Since(start))
}
