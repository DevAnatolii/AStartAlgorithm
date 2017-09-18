package main

import (
	"errors"
	"strings"
	"bytes"
)

const (
	DIRECT_MOVEMENT_PRICE   = 3
	DIAGONAL_MOVEMENT_PRICE = 4

	OBSTACLE_SYMBOL         = 'X'
	EMPTY_CELL_SYMBOL       = 'O'
	INITIAL_POSITION_SYMBOL = 'A'
	TARGET_POSITION_SYMBOL  = 'B'
	PATH_SYMBOL             = '*'

	OBSTACLE_INT_VALUE         = -1
	EMPTY_CELL_INT_VALUE       = 0
	INITIAL_POSITION_INT_VALUE = 11
	TARGET_POSITION_INT_VALUE  = 22
	PATH_INT_VALUE             = 1
)

var symbolToValue = map[rune]int{
	OBSTACLE_SYMBOL:         OBSTACLE_INT_VALUE,
	EMPTY_CELL_SYMBOL:       EMPTY_CELL_INT_VALUE,
	INITIAL_POSITION_SYMBOL: INITIAL_POSITION_INT_VALUE,
	TARGET_POSITION_SYMBOL:  TARGET_POSITION_INT_VALUE,
	PATH_SYMBOL:             PATH_INT_VALUE,
}

var valueToSymbol = map[int]rune{
	OBSTACLE_INT_VALUE:         OBSTACLE_SYMBOL,
	EMPTY_CELL_INT_VALUE:       EMPTY_CELL_SYMBOL,
	INITIAL_POSITION_INT_VALUE: INITIAL_POSITION_SYMBOL,
	TARGET_POSITION_INT_VALUE:  TARGET_POSITION_SYMBOL,
	PATH_INT_VALUE:             PATH_SYMBOL,
}

func ParseAStarTable(rawText *string) *AStarTable {
	var layout [][]int
	rowCount, colCount, rawLayout := obtainLayoutInfo(rawText)

	for rowIndex := 0; rowIndex < rowCount; rowIndex++ {
		leftBound := rowIndex * colCount
		rightBound := leftBound + colCount
		layout = append(layout, rawLayout[leftBound: rightBound])
	}
	startPoint, finishPoint := searchInitialAndTarget(&layout)

	return &AStarTable{
		InitialPoint: Point(startPoint),
		TargetPoint:  Point(finishPoint),
		Layout:       layout,
		ColumnCount:  colCount,
		RowCount:     rowCount,
	}
}

func ConvertToString(table *AStarTable) string {
	var buffer bytes.Buffer
	for rowIndex, row := range table.Layout {
		for _, val := range row {
			buffer.WriteRune(valueToSymbol[val])
		}

		if rowIndex < len(table.Layout)-1 {
			buffer.WriteString("\n")
		}
	}

	return buffer.String()
}

func obtainLayoutInfo(rawText *string) (int, int, []int) {
	rows := strings.Split(*rawText, "\n")
	rowCount, colCount := len(rows), len(rows[1])
	rowLayout := make([]int, rowCount*colCount)

	for rowIndex, row := range rows {
		for colIndex, symbol := range []rune(row) {
			if intValue, ok := symbolToValue[symbol]; ok {
				rowLayout[rowIndex*colCount+colIndex] = intValue
			} else {
				panic(errors.New("invalid symbol"))
			}
		}
	}

	return rowCount, colCount, rowLayout
}

func searchInitialAndTarget(layout *[][]int) (Point, Point) {
	var initialPosition, targetPosition = Point{-1, -1}, Point{-1, -1}

	for rowIndex, rowData := range *layout {
		for colIndex, value := range rowData {
			if value == INITIAL_POSITION_INT_VALUE {
				initialPosition = Point{rowIndex, colIndex}
			}

			if value == TARGET_POSITION_INT_VALUE {
				targetPosition = Point{rowIndex, colIndex}
			}
		}
	}

	if initialPosition.RowIndex == -1 || targetPosition.ColIndex == -1 {
		panic(errors.New("either start position or finish position is not set"))
	}
	return initialPosition, targetPosition
}
