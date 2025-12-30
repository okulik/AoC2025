package day6

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day6/input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = file.Close() }()

	fmt.Printf("Corrected homework sum is %d\n", CorrectedCephalopodHomeworkSum(file))
}

func CorrectedCephalopodHomeworkSum(input io.ReadSeeker) int64 {
	cols, ops := parseInput(input)

	var totalSum int64
	for c := range cols {
		switch ops[c].op {
		case '*':
			totalSum += mul(cols[c])
		case '+':
			totalSum += sum(cols[c])
		default:
			panic("unknown operand")
		}
	}

	return totalSum
}

func parseInput(input io.ReadSeeker) ([][]string, []*Operation) {
	ops := readOps(input)
	input.Seek(0, 0)
	cols := readCols(input, ops)

	return cols, ops
}

func sum(col []string) int64 {
	var s int64
	for _, c := range col {
		num, err := strconv.Atoi(c)
		if err != nil {
			panic("bad input")
		}
		s += int64(num)
	}

	return s
}

func mul(col []string) int64 {
	var m int64 = 1
	for _, c := range col {
		num, err := strconv.Atoi(c)
		if err != nil {
			panic("bad input")
		}
		m *= int64(num)
	}
	return m
}

type Operation struct {
	start, stop int
	op          rune
}

func readOps(input io.Reader) []*Operation {
	scanner := bufio.NewScanner(input)
	var line string
	for scanner.Scan() {
		line = scanner.Text() // keep the last line
	}

	var ops []*Operation
	var current *Operation

	for i, ch := range line {
		if ch != '*' && ch != '+' {
			continue
		}

		if current != nil {
			current.stop = i - 1
			ops = append(ops, current)
		}
		current = &Operation{start: i, op: ch}
	}

	if current != nil {
		current.stop = -1
		ops = append(ops, current)
	}

	return ops
}

func readCols(input io.Reader, ops []*Operation) [][]string {
	cols := make([][]string, len(ops))

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}

		for i, op := range ops {
			start := op.start
			stop := op.stop
			if stop < 0 {
				stop = len(line)
			}

			if cols[i] == nil {
				cols[i] = make([]string, stop-start)
			}

			for j := start; j < stop; j++ {
				b := line[j]
				if b >= '0' && b <= '9' {
					cols[i][j-start] += string(b)
				}
			}
		}
	}

	return cols
}
