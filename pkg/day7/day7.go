package day7

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day7/input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = file.Close() }()

	fmt.Printf("Total beam split count is %d\n", TotalSplitCount(file))
	file.Seek(0, 0)
	fmt.Printf("Total quantum timelines count id %d\n", TotalQuantumTimelinesCount(file))
}

// TotalSplitCount calculates the total number of beam splits in the grid.
// It processes the grid row by row, following beam propagation rules:
// - '|' and 'S' symbols represent active beams that can split
// - '.' represents empty space that can be filled by a beam
// - '^' represents a beam splitter that creates beams to the left and right
// Returns the total count of beam splitter activations.
func TotalSplitCount(input io.ReadSeeker) int64 {
	grid := parseInput(input)

	var counter int64
	for r := 1; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			// Only process cells that have an active beam above them
			if grid[r-1][c] != '|' && grid[r-1][c] != 'S' {
				continue
			}

			switch grid[r][c] {
			case '.':
				// Empty space gets filled by vertical beam
				grid[r][c] = '|'
			case '^':
				// Beam splitter creates horizontal beams
				if c-1 >= 0 && grid[r][c-1] == '.' {
					grid[r][c-1] = '|'
				}
				if c+1 < len(grid[r]) && grid[r][c+1] == '.' {
					grid[r][c+1] = '|'
				}
				counter++
			}
		}
	}

	return counter
}

// TotalQuantumTimelinesCount calculates the total number of quantum timelines
// This function tracks the number of ways beams can propagate through the grid,
// counting all possible quantum paths. It uses dynamic programming to count
// the number of timelines reaching each cell. Returns the total count of all
// quantum timelines.
func TotalQuantumTimelinesCount(input io.ReadSeeker) int64 {
	grid := parseInput(input)

	// Initialize counter grid with the same dimensions as the input grid
	gridCounter := make([][]int64, len(grid))
	for r := range grid {
		gridCounter[r] = make([]int64, len(grid[r]))
		if r != 0 {
			continue
		}
		for c := range gridCounter[r] {
			// Find the starting point 'S' and initialize it with 1 timeline
			if grid[r][c] == 'S' {
				gridCounter[r][c] = 1
				break
			}
		}
	}

	var sum int64
	for r := 1; r < len(grid); r++ {
		sum = 0
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == '.' {
				// Empty space inherits timelines from the cell above
				gridCounter[r][c] += gridCounter[r-1][c]
			}

			if grid[r][c] == '^' {
				// Beam splitter distributes timelines horizontally
				if c-1 >= 0 && grid[r-1][c] == '.' {
					gridCounter[r][c-1] += gridCounter[r-1][c]
				}
				if c+1 < len(gridCounter[r]) && grid[r-1][c] == '.' {
					gridCounter[r][c+1] += gridCounter[r-1][c]
				}
			}

			sum += gridCounter[r][c]
		}
	}

	return sum
}

func parseInput(input io.ReadSeeker) [][]rune {
	grid := make([][]rune, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		grid = append(grid, []rune(line))
	}

	return grid
}
