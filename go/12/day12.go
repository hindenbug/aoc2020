package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var directions = []rune{'E', 'S', 'W', 'N'}

type Instruction struct {
	Direction rune
	Value     int
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	instructions := make([]Instruction, 0)

	for scanner.Scan() {
		instruction := strings.TrimSpace(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}
		value, _ := strconv.Atoi(instruction[1:])

		instructions = append(instructions, Instruction{rune(instruction[0]), value})
	}

	fmt.Println(manhattanDistance(instructions))
}

func manhattanDistance(instructions []Instruction) int {
	var heading, posX, posY int

	for _, instruction := range instructions {
		switch instruction.Direction {
		case 'N':
			posY += instruction.Value
		case 'S':
			posY -= instruction.Value
		case 'E':
			posX += instruction.Value
		case 'W':
			posX -= instruction.Value
		case 'L':
			heading = (heading - instruction.Value/90 + 4) % 4
		case 'R':
			heading = (heading + instruction.Value/90) % 4
		case 'F':
			switch directions[heading] {
			case 'E':
				posX += instruction.Value
			case 'W':
				posX -= instruction.Value
			case 'N':
				posY += instruction.Value
			case 'S':
				posY -= instruction.Value
			}
		}
	}

	return Abs(posX) + Abs(posY)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
