package main

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part2() int {
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		panic(err)
	}

	str := string(file)

	elves := strings.Split(str, "\n\n")

	calorieCounts := make([]int, 0)

	for i := 0; i < len(elves); i++ {
		cals := strings.Split(elves[i], "\n")

		total := 0

		for k := 0; k < len(cals); k++ {
			cal, err := strconv.Atoi(cals[k])

			if err == nil {
				total += cal
			}
		}

		calorieCounts = append(calorieCounts, total)
	}

	sort.Ints(calorieCounts)

	topThreeTotal := 0

	for i := len(calorieCounts) - 3; i < len(calorieCounts); i++ {
		topThreeTotal += calorieCounts[i]
	}

	return topThreeTotal
}
