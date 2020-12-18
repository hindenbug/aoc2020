package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	numbers := strings.Split(string(data), ",")
	gameMap, last := prepData(numbers)
	fmt.Println(solve(gameMap, last, 2020))

	gameMap, last = prepData(numbers)
	fmt.Println(solve(gameMap, last, 30000000))

}

func prepData(numbers []string) (map[int]int, int) {
	gameMap := make(map[int]int)

	for turn, val := range numbers[:len(numbers)-1] {
		number, _ := strconv.Atoi(val)
		gameMap[number] = turn + 1
	}

	last, _ := strconv.Atoi(numbers[len(numbers)-1])

	return gameMap, last
}

func solve(gameMap map[int]int, last int, cycles int) int {
	var next int

	for turn := len(gameMap) + 1; turn < cycles; turn++ {
		if val, ok := gameMap[last]; ok {
			next = turn - val
		} else {
			next = 0
		}

		gameMap[last] = turn
		last = next
	}

	return next
}
