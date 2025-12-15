package day3_test

import (
	"fmt"
	"strings"

	day3 "github.com/okulik/AoC2025/pkg/day3"
)

var testInput = `987654321111111
811111111111119
234234234234278
818181911112111`

func ExampleTotalJoltage() {
	sum := day3.TotalJoltage(strings.NewReader(testInput))
	fmt.Println(sum)
	// Output:
	// 357
}

func ExampleTotalJoltage12() {
	sum := day3.TotalJoltage12(strings.NewReader(testInput))
	fmt.Println(sum)
	// Output:
	// 3121910778619
}
