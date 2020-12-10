package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("day10_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lst := []int{0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		lst = append(lst, n)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
	}

	sort.Ints(lst)

	var diffs [3]int
	for i := 1; i < len(lst); i++ {
		diffs[lst[i]-lst[i-1]-1]++
	}
	diffs[2] += 1
	fmt.Println("Day 10 part 1:", diffs[0]*diffs[2])

	dp := make([]int, len(lst))
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		j := i - 1
		for j >= 0 && lst[i]-lst[j] <= 3 {
			dp[i] += dp[j]
			j--
		}
	}
	fmt.Println("Day 10 part 2:", dp[len(dp)-1])
}
