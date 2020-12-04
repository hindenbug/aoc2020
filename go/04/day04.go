package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	validPassports := 0

	if err != nil {
		log.Fatal(err)
	}

	passports := strings.Split(string(data), "\n\n")

	for _, passport := range passports {
		if passportHasFields(passport) {
			validPassports++
		}
	}

	fmt.Println(validPassports)
}

func passportHasFields(passport string) bool {
	for _, field := range validFields {
		if !strings.Contains(passport, field) {
			return false
		}
	}
	return true
}
