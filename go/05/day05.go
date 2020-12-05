package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")

	var highestSeatID, lowestSeatID int

	if err != nil {
		log.Fatal(err)
	}

	boardingPasses := strings.Split(string(data), "\n")

	var seatIDs [128 * 8]bool

	for _, boardingPass := range boardingPasses {
		seatID := findSeatID(boardingPass)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
		if seatID < lowestSeatID {
			lowestSeatID = seatID
		}

		seatIDs[seatID] = true
	}

	var mySeat int
	for i, taken := range seatIDs[lowestSeatID:highestSeatID] {
		if !taken {
			mySeat = lowestSeatID + i
		}
	}

	fmt.Println(highestSeatID)
	fmt.Println(mySeat)
}

func findSeatID(input string) int {
	row, col := search(127, input[:7]), search(7, input[7:])
	return (row * 8) + col
}

func search(max int, code string) int {
	lower := 0
	upper := max

	for i, char := range code {
		if string(char) == "F" || string(char) == "L" {
			upper -= int(math.Ceil(float64(upper-lower) / 2))
			if i == len(code)-1 {
				return lower
			}
		}

		if string(char) == "B" || string(char) == "R" {
			lower += int(math.Ceil(float64(upper-lower) / 2))
			if i == len(code)-1 {
				return upper
			}
		}
	}

	return 0
}
