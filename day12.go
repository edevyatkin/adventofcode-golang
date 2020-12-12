package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("day12_input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data := make([]string, 0)
	for scanner := bufio.NewScanner(file); scanner.Scan(); data = append(data, scanner.Text()) {
	}
	x, y := 0, 0
	rotate := map[string]string{"N": "NESW", "E": "ESWN", "S": "SWNE", "W": "WNES"}
	var curdir = "E"
	for _, line := range data {
		direction := string(line[0])
		value, _ := strconv.Atoi(line[1:])
		switch direction {
		case "L":
			value = 360 - value
			fallthrough
		case "R":
			curdir = string(rotate[curdir][value/90])
		case "F":
			direction = curdir
		}
		switch direction {
		case "N":
			x += value
		case "E":
			y += value
		case "S":
			x -= value
		case "W":
			y -= value
		}
	}
	println("Day 12 part 1:", int(math.Abs(float64(x))+math.Abs(float64(y))))

	x, y = 0, 0
	wpx, wpy := 10, 1
	for _, line := range data {
		direction := string(line[0])
		value, _ := strconv.Atoi(line[1:])
		switch direction {
		case "L":
			value = 360 - value
			fallthrough
		case "R":
			if value == 90 {
				wpx, wpy = wpy, -wpx
			} else if value == 180 {
				wpx, wpy = -wpx, -wpy
			} else if value == 270 {
				wpx, wpy = -wpy, wpx
			}
		case "F":
			x += wpx * value
			y += wpy * value
		}
		switch direction {
		case "N":
			wpy += value
		case "E":
			wpx += value
		case "S":
			wpy -= value
		case "W":
			wpx -= value
		}
	}
	println("Day 12 part 2:", int(math.Abs(float64(x))+math.Abs(float64(y))))
}
