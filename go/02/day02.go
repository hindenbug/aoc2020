package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	inputs := []string{}

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	passwordPolicies := processPasswordPolicies(inputs)
	fmt.Println(getValidPasswords(passwordPolicies))

}

type PasswordPolicy struct {
	min      int
	max      int
	char     string
	password string
}

func getValidPasswords(inputs []PasswordPolicy) int {
	count := 0

	for _, policy := range inputs {
		occurence := strings.Count(policy.password, policy.char)
		if occurence >= policy.min && occurence <= policy.max {
			count++
		}
	}

	return count
}

func processPasswordPolicies(input []string) []PasswordPolicy {
	results := []PasswordPolicy{}

	for _, row := range input {
		re := regexp.MustCompile(`[- :]+`)
		passwordDetails := re.Split(row, -1)

		minOccur, err := strconv.Atoi(passwordDetails[0])
		if err != nil {
			log.Fatal(err)
		}

		maxOccur, err := strconv.Atoi(passwordDetails[1])
		if err != nil {
			log.Fatal(err)
		}

		results = append(
			results,
			PasswordPolicy{minOccur, maxOccur, passwordDetails[2], passwordDetails[3]},
		)
	}

	return results
}
