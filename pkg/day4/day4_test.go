package day4_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2025/pkg/day4"
)

var testInput = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func ExampleForkliftAcessibleRolls() {
	sum := day4.ForkliftAcessibleRolls(strings.NewReader(testInput))
	fmt.Println(sum)
	// Output:
	// 13
}

func ExampleForkliftAcessibleRollsRepeated() {
	sum := day4.ForkliftAcessibleRollsRepeated(strings.NewReader(testInput))
	fmt.Println(sum)
	// Output:
	// 43
}
