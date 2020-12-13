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
	busIDs := strings.Split(inputs[1], ",")
	earliestBusID := -1
	earliestLeave := math.MaxInt64

	ids := make(map[int]int)

	for index, busID := range busIDs {
		if busID == "x" {
			continue
		}

		interval, _ := strconv.Atoi(busID)
		ids[interval] = index
		leaveTime := int(math.Ceil(float64(timestamp)/float64(interval))) * interval

		if leaveTime < earliestLeave {
			earliestLeave = leaveTime
			earliestBusID = interval
		}
	}

	fmt.Println((earliestLeave - timestamp) * earliestBusID)

	timestamp2 := 0
	product := 1
	for k, v := range ids {
		for (timestamp2+v)%k != 0 {
			timestamp2 += product
		}
		product *= k
	}

	fmt.Println(timestamp2)
}
