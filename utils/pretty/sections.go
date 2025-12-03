package pretty

import (
	"fmt"
	"strings"
)

const width = 25

var (
	doubleLine = strings.Repeat("=", width)
	singleLine = strings.Repeat("-", width)
)

func Day(day string) {
	fmt.Println(doubleLine)
	dayStr := fmt.Sprintf("Day %s!", day)
	fmt.Println(dayStr)
	fmt.Println(doubleLine)
}

func Footer() {
	fmt.Println(doubleLine)
}

func PartOne() {
	fmt.Println("Part One :")
}

func PartTwo() {
	fmt.Println("Part Two :")
}

func Answer(answer any) {
	fmt.Printf("> %v\n", answer)
	fmt.Println(singleLine)
}
