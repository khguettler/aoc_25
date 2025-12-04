package models

import (
	"fmt"
	"iter"
)

type Grid[T any] [][]T

func (g Grid[T]) SurroundingCoords(y, x int) iter.Seq[[2]int] {
	return func(yield func([2]int) bool) {
		for _, ortho := range orthos {
			y_1 := y + ortho[0]
			x_1 := x + ortho[1]

			if !g.coordIsValid(y_1, x_1) {
				y_1 = -1
				x_1 = -1
			}

			if !yield([2]int{y_1, x_1}) {
				return
			}
		}
	}
}

var orthos = [8][2]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func (g Grid[T]) coordIsValid(y, x int) bool {
	if y < 0 || y > len(g)-1 {
		return false
	}

	if x < 0 || x > len(g[y])-1 {
		return false
	}

	return true
}

func (g Grid[T]) Get(y, x int) T {
	var out T

	if y < 0 || x < 0 {
		return out
	}

	if y > len(g)-1 {
		return out
	}

	if x > len(g[y]) {
		return out
	}

	return g[y][x]
}

func (g Grid[T]) Size() (int, int) {
	return g.Height(), g.Width()
}

func (g Grid[T]) Height() int {
	return len(g)
}

func (g Grid[T]) Width() int {
	if len(g) == 0 {
		return -1
	}

	return len(g[0])
}

func (g Grid[T]) Print() {
	for _, row := range g {
		for _, cell := range row {
			fmt.Printf("%v", cell)
		}
		fmt.Println()
	}
}
