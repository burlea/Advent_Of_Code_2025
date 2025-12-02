package day02

import (
	"burlea/Advent_of_Code_2025/utils"
	"regexp"
	"strconv"
	"strings"
)

func SolvePuzzle1(input []byte) int {
	lines := utils.ParseInputCommaSeparatedArray(input)

	var totalInvalid int

	for _, line := range lines {
		ids := strings.Split(line, "-")

		id1, _ := strconv.Atoi(ids[0])
		id2, _ := strconv.Atoi(ids[1])

		for id := id1; id <= id2; id++ {
			if invalidId2Times(strconv.Itoa(id)) {
				totalInvalid += id
			}
		}

	}
	return totalInvalid
}

func invalidId2Times(id string) bool {
	if (len(id) % 2) != 0 {
		return false
	} else {
		middleIndex := len(id) / 2
		return id[:middleIndex] == id[middleIndex:]
	}
}

func SolvePuzzle2(input []byte) int {
	lines := utils.ParseInputCommaSeparatedArray(input)

	var totalInvalid int

	for _, line := range lines {
		ids := strings.Split(line, "-")

		id1, _ := strconv.Atoi(ids[0])
		id2, _ := strconv.Atoi(ids[1])

		for id := id1; id <= id2; id++ {
			if invalidIdNTimes(strconv.Itoa(id)) {
				totalInvalid += id
			}
		}

	}
	return totalInvalid
}

func invalidIdNTimes(id string) bool {

	for i := 0; i < (len(id)/2)+1; i++ {
		subString := id[0:i]

		reGex := "^(" + regexp.QuoteMeta(subString) + ")+$"
		matched, _ := regexp.MatchString(reGex, id)
		if matched {
			return true
		}
	}

	return false
}
