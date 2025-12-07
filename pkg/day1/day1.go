package day1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day1/input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = file.Close() }()

	password, err := CalculatePassword(file, 0, 99, 50)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calculating password: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Password is: %d\n", password)
}

// Instruction represents a single movement instruction for the safe's dial.
type Instruction struct {
	Direction int // 1 for clockwise (R), -1 for counter-clockwise (L)
	Distance  int
}

// Dial represents the circular dial mechanism of the safe.
type Dial struct {
	Min       int // The minimum value on the dial (inclusive).
	Max       int // The maximum value on the dial (inclusive).
	Position  int // The current position of the dial.
	RangeSize int // The total number of positions on the dial.
}

// CalculatePassword calculates the final "password" based on the input instructions.
// It initializes a Dial and processes all instructions from the input.
func CalculatePassword(input io.Reader, min, max, startPos int) (int, error) {
	dial := NewDial(min, max, startPos)
	return dial.processInstructions(input)
}

// NewDial creates and returns a new Dial.
func NewDial(min, max, startPos int) *Dial {
	return &Dial{
		Min:       min,
		Max:       max,
		Position:  startPos,
		RangeSize: max - min + 1,
	}
}

// processInstructions reads instructions from the input and calculates the total
// number of times the dial's zero point is crossed.
func (d *Dial) processInstructions(input io.Reader) (int, error) {
	instructions, err := parseInstructions(input)
	if err != nil {
		return 0, fmt.Errorf("could not parse instructions: %w", err)
	}

	totalZeroCrossings := 0
	for _, instruction := range instructions {
		displacement := instruction.Direction * instruction.Distance

		startPos := d.Position
		endPos := d.Position + displacement

		var crossings int
		if displacement > 0 {
			// The number of times we cross the zero-point (d.Min) is the difference
			// in how many full dial rotations have been completed between the end
			// and start positions.
			crossings = calcRotations(endPos, d.Min, d.RangeSize) - calcRotations(startPos, d.Min, d.RangeSize)
		} else {
			// For anti-clockwise movement, we are interested in crossings over
			// zero-point from start-1 down to end-1.
			crossings = calcRotations(startPos-1, d.Min, d.RangeSize) - calcRotations(endPos-1, d.Min, d.RangeSize)
		}

		totalZeroCrossings += crossings
		d.Position = normalize(endPos, d.Min, d.RangeSize)
	}

	return totalZeroCrossings, nil
}

// parseInstructions reads an io.Reader and converts each line into a movement instruction.
func parseInstructions(input io.Reader) ([]Instruction, error) {
	var instructions []Instruction
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		direction := -1
		if line[0] == 'R' {
			direction = 1
		}

		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("invalid distance in line %q: %w", line, err)
		}
		instructions = append(instructions, Instruction{direction, distance})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	return instructions, nil
}

// floorDiv calculates how many full rotations (spans) away from the zero point
// (min) a value is. It also corrects for movements in anti-clockwise direction.
func calcRotations(pos, min, span int) int {
	// Shift the value so the range starts at 0
	shiftedPos := pos - min

	q := shiftedPos / span
	if shiftedPos%span != 0 && shiftedPos < 0 {
		q--
	}

	return q
}

// normalize brings a value back within the dial's range [Min, Max].
func normalize(pos, min, span int) int {
	// Shift the position by min so the range starts at 0 (makes modulo work).
	shiftedPos := pos - min

	r := shiftedPos % span
	if r < 0 {
		r += span
	}

	// Unshift the remainder by min.
	return r + min
}
