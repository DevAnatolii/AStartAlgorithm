package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

const FILE_PATH_FLAG = "path"

func ReadAStarTable() (string, bool) {
	filePath := flag.String(FILE_PATH_FLAG, "", "path to file")
	flag.Parse()
	if len(*filePath) == 0 {
		fmt.Println("File with data is not specified")
		return "", false
	}

	buf, err := ioutil.ReadFile(*filePath)
	if err != nil {
		fmt.Printf("Cannot read file %v", *filePath)
		return "", false
	} else {
		return string(buf), true
	}
}

func WriteAStarTable(aStarTable *AStarTable) {
	algorithmResult := ConvertToString(aStarTable)
	fmt.Printf("Algorithm result:\n%s", algorithmResult)
}
