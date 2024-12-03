package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines = []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var result1 int
	var result2 int
	for _, line := range lines {
		valuesStr := strings.Fields(line)
		values := make([]int, len(valuesStr))
		for i, vs := range valuesStr {
			values[i], _ = strconv.Atoi(vs)
		}

		// part 1
		diffs := []int{}
		for i := 1; i < len(values); i++ {
			diffs = append(diffs, values[i]-values[i-1])
		}
		if isSafeDiffs(diffs) {
			result1++
		}

		// part 2
		for i := 0; i < len(values); i++ {
			diffs, temp := []int{}, []int{}
			temp = append(temp, values[:i]...)
			temp = append(temp, values[i+1:]...)
			for i := 1; i < len(temp); i++ {
				diffs = append(diffs, temp[i]-temp[i-1])
			}
			if isSafeDiffs(diffs) {
				result2++
				break
			}
		}
	}

	fmt.Println(result1)
	fmt.Println(result2)
}

func isSafeDiffs(diffs []int) bool {
	positive := diffs[0] > 0
	for _, d := range diffs {
		if positive {
			if d <= 0 || d > 3 {
				return false
			}
		} else {
			if d >= 0 || d < -3 {
				return false
			}
		}
	}
	return true
}
