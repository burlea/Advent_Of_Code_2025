package day03

import (
	"burlea/Advent_of_Code_2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func SolvePuzzle1(input []byte) int {
	banks := utils.ParseInputTo2DStringArray(input)

	var joltageTotal int

	for _, bank := range banks {

		var highestNumber int = 0
		var secondHighestNumber int = 0

		for i := 0; i < len(bank)-1; i++ {
			joltage, _ := strconv.Atoi(bank[i])
			nextJoltage, _ := strconv.Atoi(bank[i+1])

			if joltage > highestNumber {
				highestNumber = joltage
				secondHighestNumber = nextJoltage
			} else if joltage > secondHighestNumber {
				secondHighestNumber = joltage
			}
		}

		lastJoltage, _ := strconv.Atoi(bank[len(bank)-1])
		if lastJoltage > secondHighestNumber {
			secondHighestNumber = lastJoltage
		}

		joltageTotal += highestNumber*10 + secondHighestNumber
	}
	return joltageTotal
}

func SolvePuzzle2(input []byte) uint64 {
	banks := utils.ParseInputTo2DStringArray(input)

	var joltageTotal uint64

	for _, bank := range banks {

		fmt.Println("bank in start:", bank)

		var numberList = []string{"0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"}

		fmt.Println("Starting Number list:", numberList)

		for i := 0; i < len(bank); i++ {
			current := bank[i]

			fmt.Println("Current Number:", current)

			for j := 0; j < len(numberList); j++ {
				fmt.Println("Current Number List:", numberList[j])

				spacesLeftInBank := len(bank) - i - 1
				spacesLeftInNumberList := len(numberList) - j - 1

				fmt.Println("Spaces left in bank:", spacesLeftInBank)
				fmt.Println("Spaces left in number list:", spacesLeftInNumberList)
				fmt.Println("index j:", j)
				fmt.Println("index i:", i)

				if current >= numberList[j] && (spacesLeftInBank >= spacesLeftInNumberList) {
					if j == 0 {
						newList := make([]string, len(numberList))
						copy(newList, bank[i:i+12])
						numberList = newList
					} else {
						fmt.Println("numberList: ", numberList)
						fmt.Println("Bank: ", bank)
						fmt.Println("First Half: ", numberList[:j])
						fmt.Println("Second Half: ", bank[i:(i+len(numberList)-j)])
						newList := make([]string, len(numberList))
						copy(newList[:j], numberList[:j])
						copy(newList[j:], bank[i:(i+len(numberList)-j)])
						numberList = newList
					}
					fmt.Println("New Number list:", numberList)
					break
				}
			}
		}

		fmt.Println("Ending Number list:", numberList)
		joltageTotal += utils.ConvertStringToUInt64(strings.Join(numberList, ""))
	}
	return joltageTotal
}
