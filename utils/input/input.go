package input

import (
	"fmt"
	"os"
	"strings"
)

const dir = "./inputs"

type input struct {
	content []byte
}

func New(path string) *input {
	fullPath := dir + "/" + path
	content, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Printf("Could not read file with path '%s'. error: %s\n", fullPath, err)
		os.Exit(1)
	}

	return &input{
		content: content,
	}
}

func Input() *input {
	return New("input.txt")
}

func Sample() *input {
	return New("sample.txt")
}

func (inp *input) String() string {
	return string(inp.content)
}

func (inp *input) Lines() []string {
	lines := strings.Split(string(inp.content), "\n")

	return lines
}

func (inp *input) Grid() [][]string {
	lines := inp.Lines()
	grid := make([][]string, len(lines))

	for y, row := range lines {
		grid[y] = strings.Split(row, "")
	}

	return grid
}
