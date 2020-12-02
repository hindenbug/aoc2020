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
	fmt.Println(getValidPasswords1(passwordPolicies))
	fmt.Println(getValidPasswords2(passwordPolicies))

}

type PasswordPolicy struct {
	min      int
	max      int
	char     string
	password string
}

// checks occurence of char/letter should be >= min and <= max
func getValidPasswords1(inputs []PasswordPolicy) int {
	count := 0

	for _, policy := range inputs {
		occurence := strings.Count(policy.password, policy.char)
		if occurence >= policy.min && occurence <= policy.max {
			count++
		}
	}

	return count
}

// checks exactly one position has the char/letter in the policy
func getValidPasswords2(inputs []PasswordPolicy) int {
	count := 0

	for _, policy := range inputs {
		firstPosChar := policy.password[policy.min-1]
		secondPosChar := policy.password[policy.max-1]
		if (string(firstPosChar) == policy.char) != (string(secondPosChar) == policy.char) {
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
