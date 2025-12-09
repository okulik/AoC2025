package day2_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2025/pkg/day2"
)

var testInput = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func ExampleSumInvalidIds() {
	sum := day2.SumInvalidIds(strings.NewReader(testInput))
	fmt.Println(sum)
	// Output:
	// 1227775554
}

func ExampleSumSomeMoreInvalidIds() {
	sum := day2.SumSomeMoreInvalidIds(strings.NewReader(testInput))
	fmt.Println(sum)
	// Output:
	// 4174379265
}
