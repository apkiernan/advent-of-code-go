package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	str := string(file)

	elves := strings.Split(str, "\n\n")

	largestCalorieCount := 0

	for i := 0; i < len(elves); i++ {
		cals := strings.Split(elves[i], "\n")

		total := 0

		for k := 0; k < len(cals); k++ {
			cal, err := strconv.Atoi(cals[k])

			if err == nil {
				total += cal
			}
		}

		if total > largestCalorieCount {
			largestCalorieCount = total
		}
	}

	fmt.Println("PART 1")
	fmt.Println(largestCalorieCount)

	fmt.Println("PART 2")

	value := Part2()
	fmt.Println(value)
}
