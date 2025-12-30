package day6_test

import (
	"fmt"
	"strings"

	"github.com/okulik/AoC2025/pkg/day6"
)

var testInput = `
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +`

func ExampleCorrectedCephalopodHomeworkSum() {
	count := day6.CorrectedCephalopodHomeworkSum(strings.NewReader(testInput))
	fmt.Println(count)
	// Output:
	// 3263827
}
