package main

import "errors"

var table AStarTable
var openList, closedList []Node

func ApplyAlgorithm(gridTable *AStarTable) (*AStarTable, error) {
	table = *gridTable
	isFinishPointReached := func() bool {
		_, ok := findNode(table.TargetPoint, &closedList)
		return ok
	}
	openList = append(openList, Node{Coordinates: table.InitialPoint})

	for !isFinishPointReached() && len(openList) != 0 {
		lowestCostNode := remove(findLowestCostNodeIndex(), &openList)
		closedList = append(closedList, lowestCostNode)
		addNeighbors(lowestCostNode)
	}

	if len(openList) == 0 {
		return &AStarTable{}, errors.New("target is unreachable")
	} else {
		buildPath()
		return &table, nil
	}
}

func addNeighbors(center Node) {
	costFunction, costEstimationFunction := fetchCostAndCostEstimationFunctions()

	for rowDelta := -1; rowDelta <= 1; rowDelta++ {
		for colDelta := -1; colDelta <= 1; colDelta++ {
			if colDelta == 0 && rowDelta == 0 {
				continue
			}

			neighborPoint := Point{
				RowIndex: center.Coordinates.RowIndex + rowDelta,
				ColIndex: center.Coordinates.ColIndex + colDelta,
			}

			if _, ok := findNode(neighborPoint, &closedList); ok || !isWalkable(neighborPoint) {
				continue
			}

			neighborNode := calculateNewNode(neighborPoint, center, costFunction, costEstimationFunction)
			existingNodePosition, ok := findNode(neighborPoint, &openList)
			if ok {
				if existingNode := openList[existingNodePosition]; neighborNode.GScore < existingNode.GScore {
					openList[existingNodePosition] = neighborNode
				}
			} else {
				openList = append(openList, neighborNode)
			}
		}
	}
}

func buildPath() {
	targetNodeIndex, _ := findNode(table.TargetPoint, &closedList)
	parentNode := closedList[targetNodeIndex].Parent
	for ; parentNode.Coordinates != table.InitialPoint; {
		table.Layout[parentNode.Coordinates.RowIndex][parentNode.Coordinates.ColIndex] = PATH_INT_VALUE
		parentNode = parentNode.Parent
	}
}

func calculateNewNode(point Point, parent Node, cost func(Point, Point) int, costEstimation func(Point) int) Node {
	G := parent.GScore + cost(point, parent.Coordinates)
	H := costEstimation(point)

	return Node{
		Coordinates: point,
		Parent:      &parent,
		GScore:      G,
		HScore:      H,
		FScore:      G + H,
	}
}

func isWalkable(point Point) bool {
	return point.RowIndex >= 0 && point.RowIndex < table.RowCount &&
		point.ColIndex >= 0 && point.ColIndex < table.ColumnCount &&
		table.Layout[point.RowIndex][point.ColIndex] != OBSTACLE_INT_VALUE
}

func findLowestCostNodeIndex() int {
	lowestNodeIndex, lowestFScore := 0, openList[0].FScore
	for index, node := range openList {
		if node.FScore < lowestFScore {
			lowestFScore = node.FScore
			lowestNodeIndex = index
		}
	}

	return lowestNodeIndex
}

func findNode(coordinates Point, slice *[]Node) (int, bool) {
	for index, value := range *slice {
		if value.Coordinates == coordinates {
			return index, true
		}
	}
	return -1, false
}

func remove(elementIndex int, slicePtr *[]Node) Node {
	slice := *slicePtr
	removedNode := slice[elementIndex]

	slice[elementIndex] = slice[len(slice)-1]
	*slicePtr = slice[:len(slice)-1] // because it returns new slice
	return removedNode
}

func fetchCostAndCostEstimationFunctions() (func(Point, Point) int, func(Point) int) {
	costFunction := func(point1, point2 Point) int {
		if !point1.IsNeighbor(point2) {
			panic(errors.New("neighbor points are expected"))
		}
		if point1.IsOnDiagonal(point2) {
			return DIAGONAL_MOVEMENT_PRICE
		} else {
			return DIRECT_MOVEMENT_PRICE
		}
	}

	estimationFunction := func(point1 Point) int {
		rowDiff, colDiff := point1.CalculateDifferences(table.TargetPoint)
		return int(rowDiff+colDiff) * DIRECT_MOVEMENT_PRICE
	}

	return costFunction, estimationFunction
}
