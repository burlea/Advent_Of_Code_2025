package rangeValue

import (
	"burlea/Advent_of_Code_2025/utils"
	"strings"
)

type RangeValue struct {
	LowerBound int
	UpperBound int
}

func (r RangeValue) NewRangeValue(newLowerBound int, newUpperBound int) any {
	panic("unimplemented")
}

func NewRangeValue(lowerBound int, upperBound int) RangeValue {
	return RangeValue{
		LowerBound: lowerBound,
		UpperBound: upperBound,
	}
}

func ParseRangeValue(line string) RangeValue {
	rangeValues := strings.Split(line, "-")

	lowerBound := utils.ConvertStringToInt(rangeValues[0])
	upperBound := utils.ConvertStringToInt(rangeValues[1])

	return RangeValue{LowerBound: lowerBound, UpperBound: upperBound}
}

func (r RangeValue) Contains(value int) bool {
	return (value >= r.LowerBound) && (value <= r.UpperBound)
}

func ConvertToRangeArray(ranges []string) []RangeValue {
	rangeArray := make([]RangeValue, 0, len(ranges))

	for _, line := range ranges {
		rangeArray = append(rangeArray, ParseRangeValue(line))
	}

	return rangeArray
}
