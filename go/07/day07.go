package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(data), "\n")

	colorMap := make(map[string]map[string]int)

	for _, input := range inputs {
		re := regexp.MustCompile(`^\w+\s\w+\sbags\scontain*`)
		re2 := regexp.MustCompile(`\d\s\w+\s\w+\sbag`)
		containerBag := re.FindString(input)
		bags := re2.FindAllString(input, -1)

		containerBagColor := strings.ReplaceAll(strings.ReplaceAll(containerBag, " bags contain", ""), " bag contain", "")
		colorMap[containerBagColor] = make(map[string]int)

		for _, containedBagcolor := range bags {
			quantity, _ := strconv.Atoi(string(containedBagcolor[0]))
			key := strings.ReplaceAll(strings.ReplaceAll(containedBagcolor[2:], " bags", ""), " bag", "")
			colorMap[containerBagColor][key] = quantity
		}
	}

	count1 := 0
	for color := range colorMap {
		count1 += findBagByColor(colorMap, color)
	}

	fmt.Println(count1)
	fmt.Println(totalBagsIn(colorMap, "shiny gold") - 1)
}

func findBagByColor(colorMap map[string]map[string]int, BagColor string) int {
	for containedBagColor := range colorMap[BagColor] {
		if containedBagColor == "shiny gold" || findBagByColor(colorMap, containedBagColor) == 1 {
			return 1
		}
	}
	return 0
}

func totalBagsIn(colorMap map[string]map[string]int, BagColor string) int {
	total := 1
	for color, count := range colorMap[BagColor] {
		total += count * totalBagsIn(colorMap, color)
	}

	return total
}
