package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var rows = make([]int, 128)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	highestSeatID := 0

	if err != nil {
		log.Fatal(err)
	}

	boardingPasses := strings.Split(string(data), "\n")

	for _, boardingPass := range boardingPasses {
		seatID := findSeatID(boardingPass)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	fmt.Println(highestSeatID)
}

func findSeatID(boardingPass string) int {
	return getRow(boardingPass)*8 + getCol(boardingPass)
}

func getRow(code string) int {
	minRow := 0
	maxRow := 127

	substr := code[:7]

	for _, char := range substr {
		maxRow, minRow = setRange(string(char), maxRow, minRow)
	}

	if string(substr[len(substr)-1]) == "F" {
		return minRow
	} else {
		return maxRow
	}
}

func getCol(code string) int {
	minCol := 0
	maxCol := 7

	substr := code[7:]

	for _, char := range substr {
		maxCol, minCol = setRange(string(char), maxCol, minCol)
	}

	if string(substr[len(substr)-1]) == "L" {
		return minCol
	} else {
		return maxCol
	}
}

func setRange(char string, upper int, lower int) (int, int) {
	newUpper, newLower := upper, lower

	switch string(char) {
	case "F", "L":
		newUpper = (upper + lower) / 2
	case "B", "R":
		newLower = (upper+lower)/2 + 1

	}

	return newUpper, newLower
}
