package main

import "math"

type Point struct {
	RowIndex, ColIndex int
}

func (point1 *Point) IsOnDiagonal(point2 Point) bool {
	rowDifference, colDifference := point1.CalculateDifferences(point2)

	return rowDifference == colDifference
}

func (point1 *Point) IsNeighbor(point2 Point) bool {
	rowDifference, colDifference := point1.CalculateDifferences(point2)

	return rowDifference == 1 || colDifference == 1
}

func (point1 *Point) CalculateDifferences(point2 Point) (float64, float64) {
	return math.Abs(float64(point1.RowIndex - point2.RowIndex)), math.Abs(float64(point1.ColIndex - point2.ColIndex))
}

