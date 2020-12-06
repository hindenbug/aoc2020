package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	groups := strings.Split(string(data), "\n\n")
	var totalYesCount int
	var sameAnswerCount int

	for _, group := range groups {
		re := regexp.MustCompile(`\n`)
		groupAnswers := re.Split(group, -1)

		totalYesCount += allYesAnswers(groupAnswers)
		sameAnswerCount += allSameAnswers(groupAnswers)
	}

	fmt.Println(totalYesCount)
	fmt.Println(sameAnswerCount)
}

func allYesAnswers(groupAnswers []string) int {
	var characters [26]int
	var yesCount int
	for _, answer := range strings.Join(groupAnswers, "") {
		if characters[answer-97] == 0 {
			characters[answer-97]++
			yesCount++
		}
	}

	return yesCount
}

func allSameAnswers(groupAnswers []string) int {
	var characters [26]int
	var sameAnswerCount int
	for _, answer := range strings.Join(groupAnswers, "") {
		characters[answer-97]++
		if characters[answer-97] == len(groupAnswers) {
			sameAnswerCount++
		}
	}
	return sameAnswerCount
}
