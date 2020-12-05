package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var validFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	validPassports1, validPassports2 := 0, 0

	if err != nil {
		log.Fatal(err)
	}

	passports := strings.Split(string(data), "\n\n")

	for _, passport := range passports {
		if passportHasFields(passport) {
			validPassports1++
		}
	}

	for _, passport := range passports {
		if passportHasFields(passport) && passportIsValid(passport) {
			validPassports2++
		}
	}

	fmt.Println(validPassports1)

	fmt.Println(validPassports2)
}

func passportHasFields(passport string) bool {
	for _, field := range validFields {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}

func passportIsValid(passport string) bool {
	re := regexp.MustCompile(`\s|\n`)
	passportFields := re.Split(passport, -1)

	for _, field := range passportFields {
		keyValuePair := strings.Split(field, ":")
		if !isValidField(keyValuePair[0], keyValuePair[1]) {
			return false
		}
	}
	return true
}

func isValidField(key string, value string) bool {
	switch key {
	case "byr":
		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return year >= 1920 && year <= 2002
	case "iyr":
		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return year >= 2010 && year <= 2020
	case "eyr":
		year, err := strconv.Atoi(value)
		if err != nil {
			return false
		}
		return year >= 2020 && year <= 2030
	case "hgt":
		metric := value[len(value)-2:]
		if metric == "cm" || metric == "in" {
			val, _ := strconv.Atoi(value[:len(value)-2])
			return (metric == "cm" && val >= 150 && val <= 193) || (metric == "in" && val >= 59 && val <= 76)
		}

		return false
	case "hcl":
		re := regexp.MustCompile("^#[0-9a-f]{6}$")
		return re.MatchString(value)
	case "ecl":
		re := regexp.MustCompile("^amb|blu|brn|gry|grn|hzl|oth$")
		return re.MatchString(value)
	case "pid":
		re := regexp.MustCompile(`^\d{9}$`)
		return re.MatchString(value)
	case "cid":
		return true
	default:
		return false
	}
}
