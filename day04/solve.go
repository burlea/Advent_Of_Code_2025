package day04

import (
	"burlea/Advent_of_Code_2025/utils"
)

func SolvePuzzle1(input []byte) int {
	rollMap := utils.ParseInputTo2DStringArray(input)

	totalRows := len(rollMap)
	totalCols := len(rollMap[0])

	var totalAccessableRolls int

	for i := 0; i < totalRows; i++ {
		for j := 0; j < totalCols; j++ {

			if rollMap[i][j] == "." {
				continue
			}

			var totalAdjacentRolls int = 0

			// left up
			if withinBounds(i-1, j-1, totalRows, totalCols) && rollMap[i-1][j-1] == "@" {
				totalAdjacentRolls++
			}
			// left
			if withinBounds(i-1, j, totalRows, totalCols) && rollMap[i-1][j] == "@" {
				totalAdjacentRolls++
			}
			// left down
			if withinBounds(i-1, j+1, totalRows, totalCols) && rollMap[i-1][j+1] == "@" {
				totalAdjacentRolls++
			}
			// down
			if withinBounds(i, j+1, totalRows, totalCols) && rollMap[i][j+1] == "@" {
				totalAdjacentRolls++
			}
			// right down
			if withinBounds(i+1, j+1, totalRows, totalCols) && rollMap[i+1][j+1] == "@" {
				totalAdjacentRolls++
			}
			// right
			if withinBounds(i+1, j, totalRows, totalCols) && rollMap[i+1][j] == "@" {
				totalAdjacentRolls++
			}
			// right up
			if withinBounds(i+1, j-1, totalRows, totalCols) && rollMap[i+1][j-1] == "@" {
				totalAdjacentRolls++
			}
			// up
			if withinBounds(i, j-1, totalRows, totalCols) && rollMap[i][j-1] == "@" {
				totalAdjacentRolls++
			}

			if totalAdjacentRolls < 4 {
				totalAccessableRolls++
			}
		}
	}
	return totalAccessableRolls
}

func withinBounds(i int, j int, totalRows int, totalCols int) bool {
	return i >= 0 && j >= 0 && i < totalRows && j < totalCols
}

func SolvePuzzle2(input []byte) int {
	rollMap := utils.ParseInputTo2DStringArray(input)

	var totalRollsRemoved int
	var currentMap [][]string = rollMap

	for {
		removed, updatedRollMap := RemoveRolls(currentMap)
		totalRollsRemoved += removed
		currentMap = updatedRollMap

		if removed == 0 {
			break
		}
	}

	return totalRollsRemoved
}

func RemoveRolls(rollMap [][]string) (int, [][]string) {

	totalRows := len(rollMap)
	totalCols := len(rollMap[0])

	var totalAccessableRolls int

	for i := 0; i < totalRows; i++ {
		for j := 0; j < totalCols; j++ {

			if rollMap[i][j] == "." {
				continue
			}

			var totalAdjacentRolls int = 0

			// left up
			if withinBounds(i-1, j-1, totalRows, totalCols) && rollMap[i-1][j-1] == "@" {
				totalAdjacentRolls++
			}
			// left
			if withinBounds(i-1, j, totalRows, totalCols) && rollMap[i-1][j] == "@" {
				totalAdjacentRolls++
			}
			// left down
			if withinBounds(i-1, j+1, totalRows, totalCols) && rollMap[i-1][j+1] == "@" {
				totalAdjacentRolls++
			}
			// down
			if withinBounds(i, j+1, totalRows, totalCols) && rollMap[i][j+1] == "@" {
				totalAdjacentRolls++
			}
			// right down
			if withinBounds(i+1, j+1, totalRows, totalCols) && rollMap[i+1][j+1] == "@" {
				totalAdjacentRolls++
			}
			// right
			if withinBounds(i+1, j, totalRows, totalCols) && rollMap[i+1][j] == "@" {
				totalAdjacentRolls++
			}
			// right up
			if withinBounds(i+1, j-1, totalRows, totalCols) && rollMap[i+1][j-1] == "@" {
				totalAdjacentRolls++
			}
			// up
			if withinBounds(i, j-1, totalRows, totalCols) && rollMap[i][j-1] == "@" {
				totalAdjacentRolls++
			}

			if totalAdjacentRolls < 4 {
				totalAccessableRolls++
				rollMap[i][j] = "."
			}
		}
	}
	return totalAccessableRolls, rollMap
}
