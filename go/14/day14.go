package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input_test.txt")

	if err != nil {
		log.Fatal(err)
	}

	program := strings.Split(string(data), "\n")
	memMap := make(map[int]int)
	var mask string
	var sum int

	for _, line := range program {
		instructionSet := strings.Split(line, " = ")
		instruction := strings.TrimSpace(instructionSet[0])
		length := len(instruction)
		operand := strings.TrimSpace(instructionSet[1])

		if "mas" == instruction[0:3] {
			mask = operand
		} else if "mem" == instruction[0:3] {
			addr, _ := strconv.Atoi(instruction[4:(length - 1)])
			operand, _ := strconv.Atoi(operand)
			memMap[addr] = operand

			// from sample input

			// XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X

			// 11 ---> existing value at mem addr
			// mask 0
			// 000000000000000000000000000000001011 (11)
			// 000000000000000000000000000000000010 bit clear AND NOT (i = 1)
			// 000000000000000000000000000000001001 (9)
			// mask 1
			// 000000000000000000000000000001000000 (i = 2)
			// 000000000000000000000000000000001001 OR
			// 000000000000000000000000000001001001 (73) ---> new value at mem addr

			// 0 ---> existing value at mem  addrs
			// mask 0
			// 000000000000000000000000000000000000 (0)
			// 000000000000000000000000000000000010 bit clear AND NOT (i = 1)
			// 000000000000000000000000000000000000 (2)
			// mask 1
			// 000000000000000000000000000001000000 (i = 2)
			// 000000000000000000000000000000000000 OR
			// 000000000000000000000000000001000000 (64) ---> new value at mem addr

			for i := 0; i < 36; i++ {
				if mask[35-i] == '1' {
					operand |= (1 << i)
				}
				if mask[35-i] == '0' {
					operand &= ^(1 << i)
				}
			}

			memMap[addr] = operand
		}
	}

	for _, v := range memMap {
		sum += v
	}

	fmt.Println(sum)
}
