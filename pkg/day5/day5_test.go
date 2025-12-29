package day5_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2025/pkg/day5"
)

var testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func ExampleFreshIngredientsCount() {
	count := day5.FreshIngredientsCount(strings.NewReader(testInput))
	fmt.Println(count)
	// Output:
	// 3
}

func ExampleTotalFreshIngredientsCount() {
	total := day5.TotalFreshIngredientsCount(strings.NewReader(testInput))
	fmt.Println(total)
	// Output:
	// 14
}
