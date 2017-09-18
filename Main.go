package main

import "fmt"

func main() {
	rawText, ok := ReadAStarTable()

	if ok {
		aStarTable := ParseAStarTable(&rawText)
		aStarTable, err := ApplyAlgorithm(aStarTable)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			WriteAStarTable(aStarTable)
		}
	}
}
