package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	terrain := strings.Split(string(data), "\n")

	fmt.Println(findTrees(terrain, 3, 1))
	fmt.Println(findTrees(terrain, 1, 1) * findTrees(terrain, 3, 1) * findTrees(terrain, 5, 1) * findTrees(terrain, 7, 1) * findTrees(terrain, 1, 2))

}

func findTrees(terrain []string, x int, y int) int {
	trees, xCord, yCord := 0, 0, 0

	for yCord < len(terrain) {
		curr := terrain[yCord][xCord]

		if curr == '#' {
			trees++
		}
		xCord = (xCord + x) % len(terrain[0])
		yCord += y

	}

	return trees
}
