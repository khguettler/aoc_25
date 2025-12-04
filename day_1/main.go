package main

import (
	"github.com/khguettler/aoc_25/utils/conv"
	"github.com/khguettler/aoc_25/utils/input"
	"github.com/khguettler/aoc_25/utils/pretty"
)

func main() {
	pretty.Day("One")

	inp := input.Input()

	lines := inp.Lines()

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	pretty.PartOne()

	count := 0
	pos := 50

	for _, move := range lines {
		dir := string(move[0])
		num := conv.StrToInt(move[1:])

		num = num % 100

		if dir == "R" {
			pos += num
			pos = pos % 100
		} else {
			pos -= num
			if pos < 0 {
				pos = 100 + pos
			}
		}
		if pos == 0 {
			count++
		}
	}

	pretty.Answer(count)
}

func partTwo(lines []string) {
	pretty.PartTwo()

	count := 0
	pos := 50
	var prevPos int

	for _, move := range lines {
		dir := string(move[0])
		num := conv.StrToInt(move[1:])

		div := num / 100
		count += div

		num = num % 100
		prevPos = pos

		if dir == "L" {
			pos -= num
			pos = pos % 100

			if pos < 0 {
				pos = 100 + pos
				if prevPos != 0 && pos != 0 {
					count++
				}
			}
		} else {
			pos += num

			if pos > 99 {
				pos = pos % 100
				if prevPos != 0 && pos != 0 {
					count++
				}
			}
		}

		if pos == 0 {
			count++
		}
	}

	pretty.Answer(count)
}
