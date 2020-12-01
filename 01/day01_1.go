package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var numList []int

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		val, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}
		numList = append(numList, val)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var match = false
	for _, num := range numList {
		for _, nextNum := range numList {
			if num+nextNum == 2020 {
				fmt.Println(num, "+", nextNum, "=", num*nextNum)
				match = true
				break
			}
		}
		if match == true {
			break
		}
	}
}
