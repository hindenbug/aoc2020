package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	STEPS  = 6
	ACTIVE = '#'
)

type ConwayCube struct {
	X, Y, Z int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	currentGen := make(map[ConwayCube]bool)
	contents := strings.Split(string(data), "\n")

	for y, line := range contents {
		for x, char := range line {
			switch rune(char) {
			case ACTIVE:
				currentGen[ConwayCube{x, y, 0}] = true
			}
		}
	}

	for i := 0; i < STEPS; i++ {
		currentGen = tick(currentGen)
	}

	fmt.Println(len(currentGen))

}

func tick(prev map[ConwayCube]bool) map[ConwayCube]bool {
	next := make(map[ConwayCube]bool)
	neighbourCount := make(map[ConwayCube]int)

	for cube := range prev {
		activeNeighbours := cube.countActiveNeighbours(prev)

		for _, neighbour := range c.neighbours() {
			if prev[neighbour] {
				activeNeighbors++
			}
		}
		if (activeNeighbours == 2) || (activeNeighbours == 3) {
			next[cube] = true
		}
	}

	for cube, neighbours := range countActiveNeighboursofInactiveCubes(prev) {
		if neighbours == 3 {
			next[cube] = true
		}
	}

	return next
}

func (c ConwayCube) neighbours() []ConwayCube {
	neighbours := make([]ConwayCube, 0)

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}
				neighbours = append(neighbours, ConwayCube{c.X + x, c.Y + y, c.Z + z})
			}
		}
	}

	return neighbours
}

func (c ConwayCube) countActiveNeighbours(prev map[ConwayCube]bool) int {
	activeNeighbors := 0

	for _, neighbour := range c.neighbours() {
		if prev[neighbour] {
			activeNeighbors++
		}
	}

	return activeNeighbors
}

func countActiveNeighboursofInactiveCubes(prev map[ConwayCube]bool) map[ConwayCube]int {
	activeNeighbors := make(map[ConwayCube]int)

	for cube := range prev {
		for _, neighbour := range cube.neighbours() {
			if !prev[neighbour] {
				activeNeighbors[neighbour] = activeNeighbors[neighbour] + 1
			}
		}
	}

	return activeNeighbors
}
