package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var directions = []rune{'E', 'S', 'W', 'N'}

type Instruction struct {
	Direction rune
	Value     int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(data), "\n")

	timestamp, _ := strconv.Atoi(inputs[0])
	busIDs := strings.Split(strings.ReplaceAll(inputs[1], ",x", ""), ",")
	earliestBusID := 0
	waitTime := math.MaxInt64

	fmt.Println(busIDs)
	for _, busID := range busIDs {
		interval, _ := strconv.Atoi(busID)
		minWaitTime := interval - (timestamp % interval)

		if minWaitTime < waitTime {
			waitTime = minWaitTime
			earliestBusID = interval
		}
	}

	fmt.Println(waitTime * earliestBusID)
}
