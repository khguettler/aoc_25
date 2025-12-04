package main

import (
	"github.com/khguettler/aoc_25/utils/input"
	"github.com/khguettler/aoc_25/utils/models"
	"github.com/khguettler/aoc_25/utils/pretty"
)

func main() {
	pretty.Day("Four")

	inp := input.Input()

	grid := inp.Grid()

	partOne(grid)
	partTwo(grid)
}

func partOne(grid models.Grid[string]) {
	pretty.PartOne()

	count := 0

	for y, row := range grid {
		for x, pos := range row {
			if pos != "@" {
				continue
			}

			surr := 0
			for point := range grid.SurroundingCoords(y, x) {
				if point[0] == -1 || point[1] == -1 {
					continue
				}

				cell := grid[point[0]][point[1]]
				if cell == "@" {
					surr++
					if surr >= 4 {
						break
					}
				}
			}

			if surr < 4 {
				count++
			}
		}
	}

	pretty.Answer(count)
}

func partTwo(grid models.Grid[string]) {
	pretty.PartTwo()

	count := 0
	removed := 0

	for {
		for y, row := range grid {
			for x, pos := range row {
				if pos != "@" {
					continue
				}

				surr := 0
				for point := range grid.SurroundingCoords(y, x) {
					if point[0] < 0 || point[1] < 0 {
						continue
					}

					cell := grid[point[0]][point[1]]

					if cell == "@" {
						surr++
						if surr >= 4 {
							break
						}
					}
				}

				if surr < 4 {
					removed++
					grid[y][x] = "."
				}
			}
		}

		if removed == 0 {
			break
		}

		count += removed
		removed = 0
	}

	pretty.Answer(count)
}
