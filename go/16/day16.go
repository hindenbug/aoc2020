package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	min, max int
}

type Field struct {
	name   string
	ranges []Range
}

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(data), "\n\n")
	ticketFields := strings.Split(inputs[0], "\n")
	//yourTickets := strings.Split(inputs[1], "\n")[1:]
	nearbyTickets := strings.Split(inputs[2], "\n")[1:]

	fieldRegex := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	var fields []Field
	var ranges []Range

	for _, input := range ticketFields {
		matches := fieldRegex.FindStringSubmatch(input)
		if matches == nil {
			break
		}
		ranges = append(ranges, Range{toInt(matches[2]), toInt(matches[3])})
		ranges = append(ranges, Range{toInt(matches[4]), toInt(matches[5])})
		fields = append(fields, Field{string(matches[1]), ranges})
	}

	fmt.Println(part1(nearbyTickets, ranges))
}

func toInt(number string) int {
	n, _ := strconv.Atoi(number)
	return n
}

func part1(nearbyTickets []string, ranges []Range) int {
	sum := 0

	for _, values := range nearbyTickets {
		for _, fieldValue := range strings.Split(values, ",") {
			value, _ := strconv.Atoi(fieldValue)
			valid := false
			for _, r := range ranges {
				if value >= r.min && value <= r.max {
					valid = true
				}
			}
			if !valid {
				sum += value
			}
		}
	}

	return sum
}
