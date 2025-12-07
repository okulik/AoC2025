package day1_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2025/pkg/day1"
)

var testInput = `
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

func ExampleCalculatePassword() {
	num, _ := day1.CalculatePassword(strings.NewReader(testInput), 1, 100, 51)
	fmt.Println(num)
	// Output:
	// 6
}
