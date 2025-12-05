package main

import (
	"slices"
	"strings"

	"github.com/khguettler/aoc_25/utils/conv"
	"github.com/khguettler/aoc_25/utils/input"
	"github.com/khguettler/aoc_25/utils/pretty"
)

func main() {
	pretty.Day("Five")

	inp := input.Input()

	partOne(inp.String())
	partTwo(inp.String())
}

func partOne(inp string) {
	pretty.PartOne()

	ranges := getIntervals(inp)
	ids := getIDs(inp)

	fn := func(id int) bool {
		for _, rng := range ranges {
			if id >= rng[0] && id <= rng[1] {
				return true
			}
		}

		return false
	}

	count := 0

	for _, id := range ids {
		isFresh := fn(id)
		if isFresh {
			count++
		}
	}

	pretty.Answer(count)
}

func partTwo(inp string) {
	pretty.PartTwo()

	all := getIntervals(inp)
	slices.SortFunc(all, func(a, b [2]int) int {
		return a[0] - b[0]
	})

	condensed := [][2]int{}

	for _, intrvl := range all {
		low := intrvl[0]
		high := intrvl[1]

		shouldAdd := true

		for idx, cond := range condensed {
			cHigh := cond[1]

			if low > cHigh {
				continue
			}

			if high < cHigh {
				shouldAdd = false
				break
			} else {
				condensed[idx][1] = high
				shouldAdd = false
				break
			}
		}

		if shouldAdd {
			condensed = append(condensed, intrvl)
		}
	}

	count := 0

	for _, intrvl := range condensed {
		size := intrvl[1] - intrvl[0] + 1
		count += size
	}

	pretty.Answer(count)
}

func getIntervals(inp string) [][2]int {
	intervalPart := strings.Split(inp, "\n\n")[0]
	intervalStrs := strings.Split(intervalPart, "\n")

	intervals := make([][2]int, len(intervalStrs))

	for idx, str := range intervalStrs {
		boundStrs := strings.Split(str, "-")
		low := conv.StrToInt(boundStrs[0])
		high := conv.StrToInt(boundStrs[1])

		intervals[idx] = [2]int{low, high}
	}

	return intervals
}

func getIDs(inp string) []int {
	idPart := strings.Split(inp, "\n\n")[1]
	idStrs := strings.Split(idPart, "\n")

	ids := make([]int, len(idStrs))
	for idx, str := range idStrs {
		ids[idx] = conv.StrToInt(str)
	}

	return ids
}
