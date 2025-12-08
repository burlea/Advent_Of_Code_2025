package day08

import (
	"burlea/Advent_of_Code_2025/utils"
	"math"
	"slices"
	"sort"
	"strings"
)

func SolvePuzzle1(input []byte) int {

	lines := utils.ParseInputToStringArray(input)
	var points = make(map[int]point)

	for index := range lines {
		parts := strings.Split(lines[index], ",")
		edges := []edge{}
		points[index] = point{index, utils.ConvertStringToInt(parts[0]), utils.ConvertStringToInt(parts[1]), utils.ConvertStringToInt(parts[2]), edges}
	}

	allEdges := []edge{}

	for _, point1 := range points {
		for _, point2 := range points {
			if point1.id < point2.id {
				distance := distance(point1, point2)
				newEdge := edge{distance, point1, point2}
				allEdges = append(allEdges, newEdge)
			}
		}
	}

	sort.Slice(allEdges, func(i, j int) bool {
		return allEdges[i].distance < allEdges[j].distance
	})

	allEdgesTrimmed := allEdges[:1000]

	circuits := getAllCircuits(allEdgesTrimmed)

	sort.Slice(circuits, func(i, j int) bool {
		return circuits[i].length > circuits[j].length
	})

	result := circuits[0].length * circuits[1].length * circuits[2].length

	return result
}

func SolvePuzzle2(input []byte) int {

	lines := utils.ParseInputToStringArray(input)
	var points = make(map[int]point)

	for index := range lines {
		parts := strings.Split(lines[index], ",")
		edges := []edge{}
		points[index] = point{index, utils.ConvertStringToInt(parts[0]), utils.ConvertStringToInt(parts[1]), utils.ConvertStringToInt(parts[2]), edges}
	}

	allEdges := []edge{}

	for _, point1 := range points {
		for _, point2 := range points {
			if point1.id < point2.id {
				distance := distance(point1, point2)
				newEdge := edge{distance, point1, point2}
				allEdges = append(allEdges, newEdge)
			}
		}
	}

	sort.Slice(allEdges, func(i, j int) bool {
		return allEdges[i].distance < allEdges[j].distance
	})

	point1, point2 := getLastConnected(allEdges, len(points))

	return point1.x * point2.x
}

func getAllCircuits(allEdges []edge) []circuit {

	var allCircuits []circuit

	for _, edge := range allEdges {
		point1 := edge.point1
		point2 := edge.point2

		circuitsWithPoints := getCircuitIfExists(allCircuits, point1, point2)

		if len(circuitsWithPoints) != 0 {
			circuitToChange := allCircuits[circuitsWithPoints[0]]

			var pointsToAdd []point

			for _, circuit := range circuitsWithPoints[1:] {
				pointsToAdd = append(pointsToAdd, allCircuits[circuit].points...)
				allCircuits = slices.Delete(allCircuits, circuit, circuit+1)
			}

			pointsToAdd = append(pointsToAdd, point1)
			pointsToAdd = append(pointsToAdd, point2)

			allCircuits[circuitsWithPoints[0]] = addIfNeeded(circuitToChange, pointsToAdd)

		} else {
			allCircuits = append(allCircuits, circuit{2, []point{point1, point2}})
		}
	}

	return allCircuits
}

func getLastConnected(allEdges []edge, totalPoints int) (point, point) {

	var allCircuits []circuit
	var allVisitedPoints []point

	for _, edge := range allEdges {
		point1 := edge.point1
		point2 := edge.point2

		circuitsWithPoints := getCircuitIfExists(allCircuits, point1, point2)

		if len(circuitsWithPoints) != 0 {
			circuitToChange := allCircuits[circuitsWithPoints[0]]

			var pointsToAdd []point

			for _, circuit := range circuitsWithPoints[1:] {
				pointsToAdd = append(pointsToAdd, allCircuits[circuit].points...)
				allCircuits = slices.Delete(allCircuits, circuit, circuit+1)
			}

			pointsToAdd = append(pointsToAdd, point1)
			pointsToAdd = append(pointsToAdd, point2)

			allCircuits[circuitsWithPoints[0]] = addIfNeeded(circuitToChange, pointsToAdd)

		} else {
			allCircuits = append(allCircuits, circuit{2, []point{point1, point2}})
		}

		if !containsPoint(allVisitedPoints, point1) {
			allVisitedPoints = append(allVisitedPoints, point1)
		}
		if !containsPoint(allVisitedPoints, point2) {
			allVisitedPoints = append(allVisitedPoints, point2)
		}

		if len(allVisitedPoints) == totalPoints && len(allCircuits) == 1 {
			return point1, point2
		}
	}

	return point{}, point{}
}

func getCircuitIfExists(allCircuits []circuit, point1 point, point2 point) []int {

	var circuitList []int
	for index := range allCircuits {
		if containsPoint(allCircuits[index].points, point1) || containsPoint(allCircuits[index].points, point2) {
			circuitList = append(circuitList, index)
		}
	}

	return circuitList
}

func containsPoint(points []point, point point) bool {
	for _, pointInList := range points {
		if pointInList.id == point.id {
			return true
		}
	}
	return false
}

func addIfNeeded(circuit circuit, points []point) circuit {
	for _, point := range points {
		if !containsPoint(circuit.points, point) {
			circuit.length += 1
			circuit.points = append(circuit.points, point)
		}
	}

	return circuit
}

func distance(point1 point, point2 point) float64 {
	x1 := float64(point1.x)
	x2 := float64(point2.x)
	y1 := float64(point1.y)
	y2 := float64(point2.y)
	z1 := float64(point1.z)
	z2 := float64(point2.z)

	return math.Sqrt(((x1 - x2) * (x1 - x2)) + ((y1 - y2) * (y1 - y2)) + ((z1 - z2) * (z1 - z2)))
}

type point struct {
	id    int
	x     int
	y     int
	z     int
	edges []edge
}

type edge struct {
	distance float64
	point1   point
	point2   point
}

type circuit struct {
	length int
	points []point
}
