package main

import (
	"github.com/khguettler/aoc_25/utils/conv"
	"github.com/khguettler/aoc_25/utils/input"
	"github.com/khguettler/aoc_25/utils/pretty"
)

func main() {
	pretty.Day("Three")

	inp := input.Input()
	lines := inp.Lines()

	partOne(lines)
	partTwo(lines)
}

func partTwo(lines []string) {
	pretty.PartTwo()

	var sum int

	for _, bank := range lines {
		sum += processBank(bank)
	}

	pretty.Answer(sum)
}

func processBank(bank string) int {
	var sum int
	length := len(bank)

	var digs string

	var startIdx int

	for dig := 12; dig > 0; dig-- {
		var digIdx int
		var high byte

		for idx := startIdx; idx <= length-dig; idx++ {
			bat := bank[idx]

			if bat > high {
				high = bat
				digIdx = idx
			}
		}

		startIdx = digIdx + 1
		digs += string(high)
	}

	sum = conv.StrToInt(digs)

	return sum
}

func partOne(lines []string) {
	pretty.PartOne()

	var sum int

	for _, bank := range lines {
		var firstIdx int
		var first rune

		for idx, bat := range bank[:len(bank)-1] {
			if bat > first {
				first = bat
				firstIdx = idx
			}
		}

		var second rune
		for _, bat := range bank[firstIdx+1:] {
			if bat > second {
				second = bat
			}
		}

		jolt := conv.StrToInt(string(first) + string(second))
		sum += jolt
	}

	pretty.Answer(sum)
}
