package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction struct {
	operation string
	argument  int
	index     int
}

type Program struct {
	accumulator        int
	executedOperations map[int]bool
}

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(data), "\n")
	accumulator := 0
	var instructionSet []*Instruction

	for index, input := range instructions {
		instructionParts := strings.Split(input, " ")
		operation := instructionParts[0]
		argument, _ := strconv.Atoi(instructionParts[1])
		instruction := Instruction{operation, argument, index}
		instructionSet = append(instructionSet, &instruction)
	}

	accumulator, _ = findInfiniteLoop(instructionSet)
	fmt.Println(accumulator)

	accumulator, _ = fixBoot(instructionSet)
	fmt.Println(accumulator)
}

func findInfiniteLoop(instructionSet []*Instruction) (int, bool) {
	var program Program
	program.executedOperations = make(map[int]bool)
	var nextInstructionIndex int
	infiniteLoop := false

	for !infiniteLoop && nextInstructionIndex < len(instructionSet) {
		instruction := instructionSet[nextInstructionIndex]
		program.executedOperations[nextInstructionIndex] = true

		switch instruction.operation {
		case "nop":
			nextInstructionIndex++
			break
		case "acc":
			program.accumulator += instruction.argument
			nextInstructionIndex++
			break
		case "jmp":
			nextInstructionIndex += instruction.argument
			break
		}

		infiniteLoop = program.executedOperations[nextInstructionIndex]
	}

	return program.accumulator, infiniteLoop
}

func fixBoot(instructionSet []*Instruction) (int, bool) {
	for _, instruction := range instructionSet {
		operation := instruction.operation
		if operation == "nop" {
			instruction.operation = "jmp"
		} else if instruction.operation == "jmp" {
			instruction.operation = "nop"
		}

		if accumulator, infiniteLoop := findInfiniteLoop(instructionSet); !infiniteLoop {
			return accumulator, true
		}

		instruction.operation = operation
	}

	return 0, false
}
