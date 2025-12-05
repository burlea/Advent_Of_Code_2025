package day05

import (
	"burlea/Advent_of_Code_2025/rangeValue"
	"burlea/Advent_of_Code_2025/utils"
	"sort"
	"strings"
)

func SolvePuzzle1(input []byte) int {
	ranges, ingrediants := utils.ParseToTwoPartsSpaced(input)

	var totalFreshIngrediants int

	for _, ingrediant := range ingrediants {
		ingrediantValue := utils.ConvertStringToUInt64(ingrediant)

		for _, line := range ranges {
			rangeValues := strings.Split(line, "-")

			lowerBound := utils.ConvertStringToUInt64(rangeValues[0])
			upperBound := utils.ConvertStringToUInt64(rangeValues[1])

			if (ingrediantValue >= lowerBound) && (ingrediantValue <= upperBound) {
				totalFreshIngrediants++
				break
			}
		}
	}

	return totalFreshIngrediants
}

func SolvePuzzle2(input []byte) int {
	lines, _ := utils.ParseToTwoPartsSpaced(input)

	ranges := rangeValue.ConvertToRangeArray(lines)

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].LowerBound < ranges[j].LowerBound
	})
	mapOfFreshIngrediants := []rangeValue.RangeValue{}

	for _, rv := range ranges {
		var isMerged bool = false

		for index, existingRange := range mapOfFreshIngrediants {
			var newLowerBound int = existingRange.LowerBound
			var newUpperBound int = existingRange.UpperBound

			if existingRange.Contains(rv.LowerBound) {
				newLowerBound = min(existingRange.LowerBound, rv.LowerBound)
				newUpperBound = max(existingRange.UpperBound, rv.UpperBound)
				isMerged = true
			}

			if existingRange.Contains(rv.UpperBound) {
				newLowerBound = min(existingRange.LowerBound, rv.LowerBound)
				newUpperBound = max(existingRange.UpperBound, rv.UpperBound)
				isMerged = true
			}

			if isMerged {
				newRange := rangeValue.RangeValue{LowerBound: newLowerBound, UpperBound: newUpperBound}
				mapOfFreshIngrediants[index] = newRange
				break
			}
		}

		if !isMerged {
			mapOfFreshIngrediants = append(mapOfFreshIngrediants, rangeValue.RangeValue{LowerBound: rv.LowerBound, UpperBound: rv.UpperBound})
		}
	}

	var totalFreshIngrediants int = 0

	for _, line := range mapOfFreshIngrediants {
		totalFreshIngrediants += (line.UpperBound - line.LowerBound + 1)
	}

	return totalFreshIngrediants
}
