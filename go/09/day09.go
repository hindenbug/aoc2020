package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

var preambleSize = 25

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(data), "\n")

	var i, invalidNumber, encryptionWeakness int
	var numbers []int

	for _, number := range input {
		value, _ := strconv.Atoi(number)
		numbers = append(numbers, value)
	}

	for i < len(numbers) {
		preamble := numbers[i : i+preambleSize]

		if i+preambleSize > len(numbers) {
			break
		}

		nextNumber := numbers[i+preambleSize]

		if !pairExists(nextNumber, preamble) {
			invalidNumber = nextNumber
			break
		}

		i++
	}

	fmt.Println(invalidNumber)

	encryptionWeakness = findEcryptionWeakness(numbers, invalidNumber)
	fmt.Println(encryptionWeakness)

}

func pairExists(sum int, preamble []int) bool {
	tmp := make([]int, len(preamble))
	copy(tmp, preamble)
	sort.Ints(tmp)

	j := 0
	k := preambleSize - 1
	exists := false

	for j < k {
		if tmp[j]+tmp[k] == sum {
			exists = true
			break
		}

		if tmp[j]+tmp[k] < sum {
			j++
		} else {
			k--
		}
	}

	return exists
}

func findEcryptionWeakness(numbers []int, invalidNumber int) int {
	weakness := 0

	for i, n1 := range numbers {
		min, max := n1, n1
		for _, n2 := range numbers[i+1:] {
			n1 += n2
			if n2 > max {
				max = n2
			}
			if min > n2 {
				min = n2
			}
			if n1 == invalidNumber {
				weakness = min + max
				break
			}
		}
	}

	return weakness
}
