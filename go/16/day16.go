package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	field                  string
	lmin, lmax, hmin, hmax int
}

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	inputs := strings.Split(string(data), "\n\n")
	ticketFields := strings.Split(inputs[0], "\n")
	myTicket := strings.Split(inputs[1], "\n")[1]
	nearbyTickets := strings.Split(inputs[2], "\n")[1:]

	fieldRegex := regexp.MustCompile(`^(.+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	fieldRules := make(map[string]Rule)
	var rules []Rule
	var fieldNames []string

	for _, input := range ticketFields {
		matches := fieldRegex.FindStringSubmatch(input)
		if matches == nil {
			break
		}
		name := string(matches[1])
		fieldNames = append(fieldNames, name)
		rules = append(rules, Rule{name, toInt(matches[2]), toInt(matches[3]), toInt(matches[4]), toInt(matches[5])})
		fieldRules[name] = Rule{name, toInt(matches[2]), toInt(matches[3]), toInt(matches[4]), toInt(matches[5])}

	}

	sum, validTickets := part1(nearbyTickets, fieldRules)
	fmt.Println(sum)
	fmt.Println(findDepartureFieldProduct(myTicket, part2(validTickets, fieldRules)))
}

func part1(tickets []string, fieldRules map[string]Rule) (int, [][]string) {
	sum := 0
	var validTickets [][]string

	for _, ticket := range tickets {
		ticket := strings.Split(ticket, ",")
		isTicketValid := true

		for _, value := range ticket {
			val, _ := strconv.Atoi(value)
			validValue := false

			for _, r := range fieldRules {
				if (val >= r.lmin && val <= r.lmax) || (val >= r.hmin && val <= r.hmax) {
					validValue = true
				}
			}

			if !validValue {
				sum += val
				isTicketValid = false
				break
			}
		}

		if isTicketValid {
			validTickets = append(validTickets, ticket)
		}
	}

	return sum, validTickets
}

func part2(tickets [][]string, fieldRules map[string]Rule) map[int]string {
	possiblePositions := make(map[string][]int)

	for name, rule := range fieldRules {
		for i := 0; i <= len(tickets[0])-1; i++ {
			possible := true
			for _, ticket := range tickets {
				val, _ := strconv.Atoi(ticket[i])
				if !((val >= rule.lmin && val <= rule.lmax) || (val >= rule.hmin && val <= rule.hmax)) {
					possible = false
					break
				}
			}

			if possible {
				possiblePositions[name] = append(possiblePositions[name], i)
			}
		}
	}

	fieldOrder := make(map[int]string)

	fmt.Println(possiblePositions)
	for len(possiblePositions) > 0 {
		for name, rules := range possiblePositions {
			if len(rules) == 1 {
				col := rules[0]
				fieldOrder[col] = name
				delete(possiblePositions, name)
				for n, values := range possiblePositions {
					possiblePositions[n] = removeFrom(values, col)
				}
				break
			}
		}
	}

	return fieldOrder
}

func removeFrom(list []int, val int) []int {
	newList := make([]int, 0)

	for _, v := range list {
		if v != val {
			newList = append(newList, v)
		}
	}
	return newList
}

func findDepartureFieldProduct(myTicket string, fieldOrder map[int]string) int {
	product := 1
	ticket := strings.Split(myTicket, ",")

	for k, v := range fieldOrder {
		if len(v) >= 9 && v[0:9] == "departure" {
			val, _ := strconv.Atoi(ticket[k])
			product *= val
		}
	}

	return product
}

func toInt(number string) int {
	n, _ := strconv.Atoi(number)
	return n
}
