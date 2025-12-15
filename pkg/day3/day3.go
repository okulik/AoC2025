package day3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day3/input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = file.Close() }()

	fmt.Printf("Sum of joltages: %d\n", TotalJoltage12(file))
}

func TotalJoltage(input io.Reader) int64 {
	banks := parseBanks(input)
	var sum int64
	for _, b := range banks {
		sum += maxJoltage(b, 2)
	}

	return sum
}

func TotalJoltage12(input io.Reader) int64 {
	banks := parseBanks(input)
	var sum int64
	for _, b := range banks {
		sum += maxJoltage(b, 12)
	}

	return sum
}

func parseBanks(input io.Reader) []string {
	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if len(l) == 0 {
			continue
		}

		lines = append(lines, l)
	}

	return lines
}

type Battery struct {
	el, ind int
}

// maxJoltage finds the maximum joltage value by selecting the last bnum digits
// and finding the maximum value in each position from left to right
func maxJoltage(b string, bnum int) int64 {
	// Extract the last bnum digits as batteries
	batteries := extractLastBatteries(b, bnum)

	// Find maximum values for each battery position
	maxBatteries := findMaximumBatteries(b, batteries)

	// Convert battery array to final number
	return batteriesToNumber(maxBatteries)
}

// extractLastBatteries gets the last bnum digits from the string as Battery structs
func extractLastBatteries(b string, bnum int) []Battery {
	startIdx := len(b) - bnum
	batteries := make([]Battery, 0, bnum)

	for i := startIdx; i < len(b); i++ {
		batteries = append(batteries, Battery{
			el:  val(b[i]),
			ind: i,
		})
	}

	return batteries
}

// findMaximumBatteries finds the maximum value for each battery position
// by looking backwards from each position to find the largest digit
func findMaximumBatteries(b string, batteries []Battery) []Battery {
	searchStart := 0

	for i := range batteries {
		// Exit if current battery is before our search boundary
		if batteries[i].ind <= searchStart {
			break
		}

		// Find the maximum digit in the range [searchStart, batteries[i].ind]
		maxPos, maxVal := findMaxInRange(b, searchStart, batteries[i].ind)

		// If we found a larger value, update the battery
		if maxPos < batteries[i].ind {
			batteries[i].el = maxVal
			batteries[i].ind = maxPos
			searchStart = maxPos + 1
		} else {
			searchStart = batteries[i].ind + 1
		}
	}

	return batteries
}

// findMaxInRange finds the position and value of the maximum digit in the given range
func findMaxInRange(b string, start, end int) (int, int) {
	maxPos := end
	maxVal := val(b[end])

	// Search backwards from end to start (inclusive)
	for j := end - 1; j >= start; j-- {
		currentVal := val(b[j])
		if currentVal >= maxVal {
			maxPos = j
			maxVal = currentVal
		}
	}

	return maxPos, maxVal
}

// batteriesToNumber converts the battery array to the final number
func batteriesToNumber(batteries []Battery) int64 {
	return sum(batteries)
}

// val converts a digit character to its integer value
func val(digit byte) int {
	return int(digit - '0')
}

// sum converts battery array to a number by treating each battery as a digit
func sum(batteries []Battery) int64 {
	var result int64

	for i := range batteries {
		// Each battery represents a digit, positioned from left to right
		// The leftmost battery has the highest place value
		placeValue := exp(10, len(batteries)-i-1)
		result += int64(batteries[i].el) * placeValue
	}

	return result
}

// exp calculates base^exponent using integer arithmetic
func exp(base, exponent int) int64 {
	result := int64(1)

	for i := 0; i < exponent; i++ {
		result *= int64(base)
	}

	return result
}
