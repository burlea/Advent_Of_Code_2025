package day09

import (
	"burlea/Advent_of_Code_2025/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

func SolvePuzzle1(input []byte) int64 {
	lines := utils.ParseInputToStringArray(input)
	var points = make(map[int]point)

	for index := range lines {
		parts := strings.Split(lines[index], ",")
		points[index] = point{index, utils.ConvertStringToInt(parts[0]), utils.ConvertStringToInt(parts[1])}
	}

	areaList := []area{}

	for _, point1 := range points {
		for _, point2 := range points {
			if point1.id != point2.id {
				areaValue := getArea(point1, point2)
				areaList = append(areaList, area{point1, point2, areaValue})
			}
		}
	}

	sort.Slice(areaList, func(i, j int) bool {
		return areaList[i].areaValue > areaList[j].areaValue
	})
	return areaList[0].areaValue
}

func SolvePuzzle2(input []byte) int64 {
	lines := utils.ParseInputToStringArray(input)
	var points = []point{}

	for index, value := range lines {
		parts := strings.Split(value, ",")

		points = append(points, point{index, utils.ConvertStringToInt(parts[0]), utils.ConvertStringToInt(parts[1])})
	}

	areaList := []area{}

	for _, point1 := range points {
		for _, point2 := range points {
			if point1.id != point2.id {
				areaValue := getArea(point1, point2)
				areaList = append(areaList, area{point1, point2, areaValue})
			}
		}
	}

	sort.Slice(areaList, func(i, j int) bool {
		return areaList[i].areaValue > areaList[j].areaValue
	})

	pointListCopy := points
	pointListCopy = append(pointListCopy, points[0])

	edgeList := []edge{}

	for i := 0; i < len(pointListCopy)-1; i++ {
		point1 := pointListCopy[i]
		point2 := pointListCopy[i+1]

		edge := edge{point1, point2}
		edgeList = append(edgeList, edge)
	}
	var largestArea int64

	for _, areaValue := range areaList {

		point1 := areaValue.point1
		point2 := areaValue.point2

		corner1 := point{0, point1.x, point2.y}
		corner2 := point{0, point2.x, point1.y}

		fmt.Println("Area: ", areaValue)
		fmt.Println("corner1: ", corner1)
		fmt.Println("corner2: ", corner2)

		if inShape(corner1, edgeList) && inShape(corner2, edgeList) {
			largestArea = areaValue.areaValue
			fmt.Println("Valid: ", areaValue)
			break
		}
	}

	return largestArea
}

type point struct {
	id int
	x  int
	y  int
}

type area struct {
	point1    point
	point2    point
	areaValue int64
}

type edge struct {
	point1 point
	point2 point
}

func getArea(point1 point, point2 point) int64 {
	x1 := float64(point1.x)
	x2 := float64(point2.x)
	y1 := float64(point1.y)
	y2 := float64(point2.y)

	return int64((math.Abs(x1-x2) + 1) * (math.Abs(y1-y2) + 1))
}

func inShape(point point, edgeValues []edge) bool {
	fmt.Println("point: ", point)

	var left int = 0
	var right int = 0
	var up int = 0
	var down int = 0

	for _, edge := range edgeValues {
		if edge.point1.x == edge.point2.x { // ver line
			if pointWithin(point.y, edge.point1.y, edge.point2.y) {
				if edge.point1.x == point.x {
					left++
					right++
				} else if edge.point1.x > point.x {
					right++
				} else {
					left++
				}
			} else if edge.point1.x == point.x {
				if edge.point1.y > point.y && edge.point2.y > point.y {
					continue
				} else if edge.point1.y < point.y && edge.point2.y < point.y {
					continue
				}
			}
		} else if edge.point1.y == edge.point2.y { // hor line
			if pointWithin(point.x, edge.point1.x, edge.point2.x) {
				if edge.point1.y == point.y {
					up++
					down++
				} else if edge.point1.y > point.y {
					down++
				} else {
					up++
				}
			} else if edge.point1.y == point.y {
				if edge.point1.x > point.x && edge.point2.x > point.x {
					continue
				} else if edge.point1.x < point.x && edge.point2.x < point.x {
					continue
				}
			}
		}
	}

	fmt.Println("left: ", left)
	fmt.Println("right: ", right)
	fmt.Println("up: ", up)
	fmt.Println("down: ", down)

	return (left%2 == 1) && (right%2 == 1) && (down%2 == 1) && (up%2 == 1)
}

func pointWithin(value int, bound1 int, bound2 int) bool {
	upperBound := max(bound1, bound2)
	lowerBound := min(bound1, bound2)

	return value <= upperBound && value >= lowerBound
}
