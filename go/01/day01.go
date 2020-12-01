package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const Result = 2020

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var numList []int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		val, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

		if err != nil {
			log.Fatal(err)
		}
		numList = append(numList, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sumOfTwo, err := findTwoNumsWithSum(numList, Result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumOfTwo)

	sumOfThree, err := findThreeNumsWithSum(numList, Result)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sumOfThree)
}

// part 1
func findTwoNumsWithSum(input []int, sum int) (int, error) {
	seen := make(map[int]bool)

	for _, num := range input {
		expected := sum - num
		if seen[num] {
			return num * expected, nil
		}
		seen[expected] = true
	}

	return 0, errors.New("no pair found")
}

// part 2
func findThreeNumsWithSum(input []int, sum int) (int, error) {
	for i, num := range input {
		expected := sum - num
		sumOfFirstTwo, err := findTwoNumsWithSum(input[i:], expected)
		if err != nil {
			continue
		} else {
			return num * sumOfFirstTwo, nil
		}
	}

	return 0, errors.New("no three numbers found")
}
