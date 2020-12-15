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

	fmt.Println(solve(prepData(numbers), 2020))
	fmt.Println(solve(prepData(numbers), 30000000))

}

func prepData(numbers []string) map[int]int {
	gameMap := make(map[int]int)

	for turn, val := range numbers {
		number, _ := strconv.Atoi(val)
		gameMap[number] = turn + 1
	}

	return gameMap
}

func solve(gameMap map[int]int, cycles int) int {
	var nextNumber int

	for turn := len(gameMap) + 1; turn < cycles; turn++ {
		if val, ok := gameMap[nextNumber]; ok {
			gameMap[nextNumber] = turn
			nextNumber = turn - val
		} else {
			gameMap[nextNumber] = turn
			nextNumber = 0
		}
	}

	return nextNumber
}
