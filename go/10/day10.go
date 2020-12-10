package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}
	var adapters []int
	//var builtInAdapterRating int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		adapter, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		if err != nil {
			log.Fatal(err)
		}

		adapters = append(adapters, adapter)
	}

	sort.Ints(adapters)

	fmt.Println(joltDiff(adapters))

}

func joltDiff(adapters []int) int {
	var prevAdapter, oneDiffs, threeDiffs int

	for _, current := range adapters {
		diff := current - prevAdapter
		prevAdapter = current

		switch diff {
		case 1:
			oneDiffs++
		case 3:
			threeDiffs++
		}
	}

	threeDiffs++

	return oneDiffs * threeDiffs
}
