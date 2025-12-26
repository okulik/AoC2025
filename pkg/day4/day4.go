package day4

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day4/input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = file.Close() }()

	fmt.Printf("Total forklift accessible roles: %d\n", ForkliftAcessibleRollsRepeated(file))
}

func ForkliftAcessibleRolls(input io.Reader) int64 {
	return counter(parseInput(input))
}

func ForkliftAcessibleRollsRepeated(input io.Reader) int64 {
	grid := parseInput(input)

	var totalCount int64
	for {
		cnt := counter(grid)
		if cnt == 0 {
			break
		}
		totalCount += cnt
	}

	return totalCount
}

func parseInput(input io.Reader) [][]byte {
	scanner := bufio.NewScanner(input)
	var grid [][]byte
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if len(l) == 0 {
			continue
		}
		ln := make([]byte, len(l))
		for i, ch := range l {
			ln[i] = byte(ch)
		}
		grid = append(grid, ln)
	}

	return grid
}

func counter(grid [][]byte) int64 {
	var counter int64
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' && isAccessible(grid, i, j) {
				counter++
				grid[i][j] = 'x'
			}
		}
	}

	return counter
}

func isAccessible(grid [][]byte, row, col int) bool {
	adjacents := [][]int{
		{row - 1, col - 1}, //↖️
		{row, col - 1},     //⬅️
		{row + 1, col - 1}, //↙️
		{row - 1, col},     //⬆️
		{row + 1, col},     //⬇️
		{row - 1, col + 1}, //↗️
		{row, col + 1},     //➡️
		{row + 1, col + 1}, //↘️
	}

	count := 0
	for _, pos := range adjacents {
		if pos[0] >= 0 && pos[0] < len(grid) &&
			pos[1] >= 0 && pos[1] < len(grid[0]) &&
			grid[pos[0]][pos[1]] == '@' {
			count++
		}
	}

	return count < 4
}
