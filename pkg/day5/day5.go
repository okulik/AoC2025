package day5

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	file, err := os.Open("pkg/day5/input")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening input file: %v\n", err)
		os.Exit(1)
	}
	defer func() { _ = file.Close() }()

	fmt.Printf("Fresh ingredients count: %d\n", FreshIngredientsCount(file))
	file.Seek(0, 0)
	total := TotalFreshIngredientsCount(file)
	fmt.Printf("Total fresh ingredients count: %d\n", total)
}

func FreshIngredientsCount(input io.Reader) int64 {
	ranges, ingredients := parseInput(input)

	slices.SortFunc(ranges, sortRanges)

	var freshCount int64
	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if ingredient >= r.from && ingredient <= r.to {
				freshCount += 1
				break
			}
		}
	}

	return freshCount
}

func TotalFreshIngredientsCount(input io.Reader) int64 {
	ranges, _ := parseInput(input)

	slices.SortFunc(ranges, sortRanges)

	for i := 0; i < len(ranges)-1; i++ {
		for j := i + 1; j < len(ranges); j++ {
			if isOverlap(ranges[i], ranges[j]) {
				ranges[i] = mergeRange(ranges[i], ranges[j])
				ranges[j].deleted = true
			}
		}
	}

	var num int64
	for _, r := range ranges {
		if r.isDeleted() {
			continue
		}
		num += r.to - r.from + 1
	}

	return num
}

type Range struct {
	from, to int64
	deleted  bool
}

func (r Range) isDeleted() bool {
	return r.deleted
}

func parseInput(input io.Reader) ([]Range, []int64) {
	ingredients := []int64{}
	ranges := []Range{}
	foundIngredients := false

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		l := strings.TrimSpace(scanner.Text())
		if len(l) == 0 {
			foundIngredients = true
			continue
		}
		if foundIngredients {
			ingred, err := strconv.Atoi(l)
			if err != nil {
				panic("invalid input")
			}
			ingredients = append(ingredients, int64(ingred))
		} else {
			rngs := strings.Split(l, "-")
			if len(rngs) != 2 {
				panic("invalid input")
			}
			from, err := strconv.Atoi(rngs[0])
			if err != nil {
				panic("invalid input")
			}
			to, err := strconv.Atoi(rngs[1])
			if err != nil {
				panic("invalid input")
			}
			ranges = append(ranges, Range{from: int64(from), to: int64(to)})
		}
	}

	return ranges, ingredients
}

func isOverlap(a, b Range) bool {
	if a.isDeleted() || b.isDeleted() {
		return false
	}

	return a.from <= b.to && b.from <= a.to
}

func mergeRange(a, b Range) Range {
	return Range{from: min(a.from, b.from), to: max(a.to, b.to)}
}

func sortRanges(a, b Range) int {
	return cmp.Compare(a.from, b.from)
}
