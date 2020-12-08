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
	instructions       []Instruction
	accumulator        int
	executedOperations map[int]bool
}

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	instructions := strings.Split(string(data), "\n")
	var program Program
	program.executedOperations = make(map[int]bool)
	program.accumulator = 0
	var nextInstructionIndex int

	for index, input := range instructions {
		instructionParts := strings.Split(input, " ")
		operation := instructionParts[0]
		argument, _ := strconv.Atoi(instructionParts[1])
		instruction := Instruction{operation, argument, index}
		program.instructions = append(program.instructions, instruction)
	}

	for nextInstructionIndex < len(program.instructions) {
		instruction := program.instructions[nextInstructionIndex]
		if program.executedOperations[nextInstructionIndex] {
			break
		}

		program.executedOperations[nextInstructionIndex] = true

		switch instruction.operation {
		case "nop":
			nextInstructionIndex++
			continue
		case "acc":
			program.accumulator += instruction.argument
			nextInstructionIndex++
			continue
		case "jmp":
			nextInstructionIndex += instruction.argument
			continue
		}
	}

	fmt.Println(program.accumulator)
}
