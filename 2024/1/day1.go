package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var lines = []string{}
	for {
		if scanner.Scan() {
			lines = append(lines, scanner.Text())
		} else {
			break
		}
	}

	// part 1
	var list1 = []int{}
	var list2 = []int{}
	for _, line := range lines {
		values := strings.Split(line, "   ")
		v1, e1 := strconv.Atoi(values[0])
		if e1 != nil {
			log.Fatal(e1)
		}
		v2, e2 := strconv.Atoi(values[1])
		if e2 != nil {
			log.Fatal(e2)
		}
		list1 = append(list1, v1)
		list2 = append(list2, v2)
	}
	slices.Sort(list1)
	slices.Sort(list2)
	var result1 int
	for i := range len(list1) {
		result1 += int(math.Abs(float64(list1[i] - list2[i])))
	}

	// part 2
	var counter = make(map[int]int)
	for _, n := range list2 {
		counter[n]++
	}
	var result2 int
	for _, n := range list1 {
		result2 += n * counter[n]
	}

	fmt.Println(result1)
	fmt.Println(result2)
}
