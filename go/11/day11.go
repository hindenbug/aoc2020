package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Position struct {
	X, Y int
}

type state string

const (
	floor    state = "."
	empty          = "L"
	occupied       = "#"
)

func main() {
	dat, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(dat)), "\n")

	var seatsMap [][]string

	for _, line := range lines {
		seatsMap = append(seatsMap, strings.Split(line, ""))
	}

	changed := true
	for changed {
		seatsMap, changed = nextState(seatsMap)
	}

	var count int
	for _, row := range seatsMap {
		for _, seat := range row {
			if seat == occupied {
				count++
			}
		}
	}

	fmt.Println(count)
}

func nextState(seatMap [][]string) ([][]string, bool) {
	changed := false
	newSeatMap := make([][]string, 0)

	for i, row := range seatMap {
		newRow := make([]string, 0)
		for j, seat := range row {
			pos := Position{i, j}

			if seat == empty && getOccupiedNeighbours(pos, seatMap) == 0 {
				newRow = append(newRow, occupied)
				changed = true
			} else if seat == occupied && getOccupiedNeighbours(pos, seatMap) >= 4 {
				newRow = append(newRow, empty)
				changed = true
			} else {
				newRow = append(newRow, seat)
			}
		}
		newSeatMap = append(newSeatMap, newRow)
	}

	return newSeatMap, changed
}

func getOccupiedNeighbours(position Position, seatMap [][]string) int {
	total := 0

	adjacents := [][]int{{1, 0}, {1, -1}, {1, 1}, {0, 1}, {0, -1}, {-1, 0}, {-1, 1}, {-1, -1}}

	for _, value := range adjacents {
		x, y := value[0], value[1]

		if position.X+x >= 0 && position.Y+y >= 0 && position.X+x <= len(seatMap)-1 && position.Y+y <= len(seatMap[0])-1 {
			if seatMap[position.X+x][position.Y+y] == occupied {
				total++
			}
		}

	}

	return total
}
