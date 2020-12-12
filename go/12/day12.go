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

	fmt.Println(manhattanDistance1(instructions))
	fmt.Println(manhattanDistance2(instructions))

}

func manhattanDistance1(instructions []Instruction) int {
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

func manhattanDistance2(instructions []Instruction) int {
	var posX, posY int
	waypointX, waypointY := 10, 1

	for _, instruction := range instructions {
		switch instruction.Direction {
		case 'N':
			waypointY += instruction.Value
		case 'S':
			waypointY -= instruction.Value
		case 'E':
			waypointX += instruction.Value
		case 'W':
			waypointX -= instruction.Value
		case 'L':
			for i := 0; i < instruction.Value/90; i++ {
				waypointX, waypointY = -waypointY, waypointX
			}
		case 'R':
			for i := 0; i < instruction.Value/90; i++ {
				waypointX, waypointY = waypointY, -waypointX
			}
		case 'F':
			posX += instruction.Value * waypointX
			posY += instruction.Value * waypointY
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
