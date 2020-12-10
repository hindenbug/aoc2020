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
	fmt.Println(arrangemments(adapters))

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

	// diff of built in adapter which is 3 from highest rated adapter
	threeDiffs++

	return oneDiffs * threeDiffs
}

func arrangemments(adapters []int) int {
	arrangements := make(map[int]int)
	arrangements[0] = 1

	for _, adapter := range adapters {
		for i := 1; i <= 3; i++ {
			diff := adapter - i
			if _, ok := arrangements[diff]; ok {
				arrangements[adapter] += arrangements[diff]
			}
		}
	}
	return arrangements[adapters[len(adapters)-1]]
}
