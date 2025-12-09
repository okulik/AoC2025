package day2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"unicode/utf8"
)

// Run executes the day 2 puzzle solution.
// It reads the input file and calculates two different error metrics:
// 1. Numbers with identical halves (e.g., 123123)
// 2. Numbers with repeated patterns of any size (e.g., 123123, 1212, 111111)
func Run() {
	file, err := os.Open("pkg/day2/input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = file.Close() }()

	fmt.Printf("Number of errors: %d\n", SumInvalidIds(file))
	file.Seek(0, 0)
	fmt.Printf("Number of errors, take two: %d\n", SumSomeMoreInvalidIds(file))
}

// SumInvalidIds calculates the sum of all numbers in the given ranges
// that have identical first and second halves. For example, 123123
// would be considered an error because "123" == "123".
func SumInvalidIds(input io.Reader) int64 {
	ranges := parseRanges(input)
	var errors int64
	for _, r := range ranges {
		for i := r.from; i <= r.to; i++ {
			if hasError(i) {
				errors += i
			}
		}
	}

	return errors
}

// SumSomeMoreInvalidIds calculates the sum of all numbers in the given ranges
// that contain repeated patterns of any size. This includes numbers like:
// - 1212 (pattern "12" repeated twice)
// - 123123 (pattern "123" repeated twice)
// - 111111 (pattern "1" repeated six times)
func SumSomeMoreInvalidIds(input io.Reader) int64 {
	ranges := parseRanges(input)
	var errors int64
	for _, r := range ranges {
		for i := r.from; i <= r.to; i++ {
			if hasSomeMoreError(i) {
				errors += i
			}
		}
	}

	return errors
}

type FromTo struct {
	from int64
	to   int64
}

// parseRanges reads comma-separated range strings from the input
// and converts them into FromTo structs. Expected format: "1-100,200-300"
func parseRanges(input io.Reader) []FromTo {
	arr := make([]FromTo, 0)

	scanner := bufio.NewScanner(input)
	scanner.Split(scanWords)
	for scanner.Scan() {
		var from, to int64
		fmt.Sscanf(scanner.Text(), "%d-%d", &from, &to)
		arr = append(arr, FromTo{from: from, to: to})
	}

	return arr
}

// scanWords is a custom split function for bufio.Scanner that splits
// input on commas instead of whitespace. It handles UTF-8 encoding
// and returns tokens between commas, skipping leading whitespace.
func scanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isSpace(r) {
			break
		}
	}
	// Scan until comma, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if r == ',' {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

// isSpace checks if a rune is a whitespace character, including
// both ASCII whitespace and Unicode space characters.
func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

// hasError checks if a number has identical first and second halves.
// Returns true for numbers like 123123 where "123" == "123".
func hasError(num int64) bool {
	// Convert number to string.
	str := strconv.FormatInt(num, 10)
	l := len(str)

	// If the number contains odd count of digits, it has no errors.
	if l%2 == 1 {
		return false
	}

	// Simply compare the first half with the second half.
	return str[0:l/2] == str[l/2:]
}

// hasSomeMoreError checks if a number contains repeated patterns of any size.
// It tests all possible pattern lengths (orders) from 1 to half the number's length.
// Returns true for numbers like 1212 (pattern "12"), 123123 (pattern "123"),
// or 111111 (pattern "1").
func hasSomeMoreError(num int64) bool {
	// Convert number to string.
	str := strconv.FormatInt(num, 10)
	l := len(str)

	// Pre-allocate array of strings that will contain tuples.
	tuples := make([]string, len(str))

	// Total number of distinct tuple orders ranges from 1 up
	// to half the length of the original number, where order
	// represents the tuple size.
	for order := 1; order <= l/2; order++ {
		// Skip matching if size of the number isn't a multiple
		// of order.
		if l%order != 0 {
			continue
		}

		// Generate all possible tuples from number while minding
		// the current tuple size.
		tuples = generateTuples(order, str, tuples)

		// If all tuples are equal, return true.
		if !slices.ContainsFunc(tuples, func(s string) bool { return s != tuples[0] }) {
			return true
		}
	}

	return false
}

// generateTuples will generate an array of chunks of order size
// from the original string. If the number is e.g. "123123",
// iterating over orders (from 1 to 3) will produce the
// following tuple arrays:
//
//	["1","2","3","1","2","3"]
//	["12", "31", "23"]
//	["123", "123"]
//
// The function will return only the required number of tuples
// and discard the remainder of the slice.
func generateTuples(order int, num string, tuples []string) []string {
	for i := 0; i < len(num)/order; i++ {
		tuples[i] = num[i*order : (i+1)*order]
	}
	return tuples[0 : len(num)/order]
}
