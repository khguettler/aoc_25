package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/khguettler/aoc_25/utils/input"
	"github.com/khguettler/aoc_25/utils/pretty"
)

func main() {
	pretty.Day("Two")
	defer pretty.Footer()

	inp := input.Input()
	rngs := createRanges(inp.String())

	partOne(rngs)
	partTwo(rngs)
}

func partOne(rngs [][2]int) {
	pretty.PartOne()

	var sum int

	for _, bounds := range rngs {
		for id := bounds[0]; id <= bounds[1]; id++ {
			idStr := strconv.Itoa(id)

			// IDs with an odd number of digits are always valid.
			if len(idStr)%2 != 0 {
				continue
			}

			middle := len(idStr) / 2
			if idStr[:middle] == idStr[middle:] {
				sum += id
			}
		}
	}

	pretty.Answer(sum)
}

func partTwo(rngs [][2]int) {
	pretty.PartTwo()

	var sum int

	for _, bounds := range rngs {
		for id := bounds[0]; id <= bounds[1]; id++ {
			if !validID(strconv.Itoa(id)) {
				sum += id
			}
		}
	}

	pretty.Answer(sum)
}

func validID(id string) bool {
	length := len(id)

	for numSecs := 2; numSecs <= length; numSecs++ {
		if length%numSecs != 0 {
			continue
		}

		secLength := length / numSecs
		firstSec := id[:secLength]
		eq := true

		for secIdx := 1; secIdx < numSecs; secIdx++ {
			secStart := secLength * secIdx
			secEnd := secLength * (secIdx + 1)
			sec := id[secStart:secEnd]

			if sec != firstSec {
				eq = false
				break
			}
		}

		if eq {
			return false
		}
	}

	return true
}

func createRanges(idsStr string) [][2]int {
	rngs := [][2]int{}

	for idRange := range strings.SplitSeq(idsStr, ",") {
		bounds := strings.Split(idRange, "-")

		lVal, err := strconv.Atoi(bounds[0])
		if err != nil {
			fmt.Printf("failed to convert string to int: %s\n", err)
			os.Exit(1)
		}

		rVal, err := strconv.Atoi(bounds[1])
		if err != nil {
			fmt.Printf("failed to convert string to int: %s\n", err)
			os.Exit(1)
		}

		rngs = append(rngs, [2]int{lVal, rVal})
	}

	return rngs
}
