package main

type AStarTable struct {
	InitialPoint, TargetPoint Point
	RowCount, ColumnCount     int
	Layout                    [][]int
}

type Node struct {
	FScore, GScore, HScore int
	Coordinates            Point
	Parent                 *Node
}
